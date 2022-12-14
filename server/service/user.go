package service

import (
	"crm/common"
	"crm/global"
	"crm/models"
	"crm/response"
	"fmt"
	"log"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	TOKEN_MAX_EXPIRE_TIME = 7 * 24 // Token最长有效期
)

type UserService struct {
	noticeService *NoticeService
}

func NewUserService() *UserService {
	userService := UserService{
		noticeService: &NoticeService{},
	}
	return &userService
}

// 用户注册
func (u *UserService) Register(param *models.UserCreateParam) int {

	// 判断用户是否存在
	var user = models.User{}
	err := global.Db.Where("email = ?", param.Email).First(&user).Error
	if err == nil && user.Id > 0 {
		return response.ErrCodeUserHasExist
	}

	// 校验验证码是否正确
	key := fmt.Sprintf("user:%s:code", param.Email)
	code := global.Rdb.Get(ctx, key).Val()
	if code != param.Code {
		return response.ErrCodeVerityCodeInvalid
	}

	// 对密码进行加密
	password, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.ErrCodeFailed
	}

	// 创建用户
	newUser := models.User{
		Email:    param.Email,
		Password: string(password),
		Status:   1,
		Created:  time.Now().Unix(),
	}
	if err := global.Db.Create(&newUser).Error; err != nil {
		return response.ErrCodeFailed
	}

	// 新用户默认订阅免费版
	subscribe := models.Subscribe{
		Version: 1,
		Created: time.Now().Unix(),
	}
	if global.Db.Table(SUBSCRIBE).Where("uid = ?", newUser.Id).First(&models.Subscribe{}).RowsAffected == 0 {
		subscribe.Uid = newUser.Id
		subscribe.Created = time.Now().Unix()
		err := global.Db.Table(SUBSCRIBE).Create(&subscribe).Error
		if err != nil {
			return response.ErrCodeFailed
		}
	} else {
		subscribe.Updated = time.Now().Unix()
		err = global.Db.Model(&models.Subscribe{}).Where("uid = ?", newUser.Id).Updates(&subscribe).Error
		if err != nil {
			return response.ErrCodeFailed
		}
	}

	// 注册通知
	u.noticeService.Create(&models.NoticeParam{
		Content: REGISTER_NOTICE_TEMPLATE,
		Creator: newUser.Id,
	})

	return response.ErrCodeSuccess
}

// 用户登录
func (u *UserService) Login(param *models.UserLoginParam) (*models.UserInfo, int) {

	// 判断用户是否存在
	var user = models.User{}
	err := global.Db.Where("email = ?", param.Email).First(&user).Error
	if err != nil {
		return nil, response.ErrCodeUserNotExist
	}

	// 校验账号密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(param.Password))
	if err != nil {
		return nil, response.ErrCOdeUserEmailOrPass
	}

	// 生成并保存Token
	token, err := common.GenToken(user.Id)
	if err != nil {
		log.Printf("[error]Login:GenerateToken:%s", err)
		return nil, response.ErrCodeFailed
	}

	userInfo := models.UserInfo{
		Uid:   user.Id,
		Token: token,
	}

	// 登录通知
	u.noticeService.Create(&models.NoticeParam{
		Content: LOGIN_NOTICE_TEMPLATE,
		Creator: userInfo.Uid,
	})

	return &userInfo, response.ErrCodeSuccess
}

// 获取验证码
func (u *UserService) GetVerifyCode(email string) int {
	// 生成6位随机数
	code := common.RandInt(100000, 999998)

	// 保存验证码
	key := fmt.Sprintf("user:%s:code", email)
	if err := global.Rdb.SetEx(ctx, key, strconv.Itoa(code), 10*time.Minute).Err(); err != nil {
		return response.ErrCodeFailed
	}

	// 发送验证码
	content := fmt.Sprintf("验证码%v，您正在找回密码，切勿向他人泄露。", code)
	err := common.SendMail(email, content)
	if err != nil {
		return response.ErrCodeVerityCodeSendFailed
	}
	return response.ErrCodeSuccess
}

// 忘记密码
func (u *UserService) ForgotPass(param *models.UserPassParam) int {
	// 判断用户是否存在
	var user = models.User{}
	if err := global.Db.Where("email = ?", param.Email).First(&user).Error; err != nil {
		return response.ErrCodeUserNotExist
	}

	// 校验验证码是否正确
	key := fmt.Sprintf("user:%s:code", param.Email)
	code := global.Rdb.Get(ctx, key).Val()
	if code != param.Code {
		return response.ErrCodeVerityCodeInvalid
	}

	// 对密码进行加密，更新密码
	password, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.ErrCodeFailed
	}
	upass := models.User{
		Password: string(password),
		Updated:  time.Now().Unix(),
	}
	err = global.Db.Model(&models.User{}).Where("email = ?", param.Email).Updates(&upass).Error
	if err != nil {
		return response.ErrCodeUserPassResetFailed
	}
	return response.ErrCodeSuccess
}

// 注销账号
func (u *UserService) Delete(param models.UserDeleteParam) int {
	// 校验验证码是否正确
	key := fmt.Sprintf("user:%s:code", param.Email)
	code := global.Rdb.Get(ctx, key).Val()
	if code != param.Code {
		return response.ErrCodeVerityCodeInvalid
	}
	err := global.Db.Transaction(func(tx *gorm.DB) error {
		mw := map[interface{}]string{
			&models.Product{}:   "creator = ?",
			&models.Customer{}:  "creator = ?",
			&models.Contract{}:  "creator = ?",
			&models.Subscribe{}: "uid = ?",
			&models.User{}:      "id = ?",
		}
		for k, v := range mw {
			if err := tx.Where(v, param.Id).Delete(k).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 获取用户信息
func (u *UserService) GetInfo(uid int64) (*models.UserPersonInfo, int) {
	var user models.UserPersonInfo
	err := global.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(USER).Where("id = ?", uid).First(&user).Error; err != nil {
			return err
		}
		var subscribe models.Subscribe
		if err := tx.Table(SUBSCRIBE).Select("version").Where("uid = ?", uid).First(&subscribe).Error; err != nil {
			return err
		}
		user.Version = subscribe.Version
		return nil
	})
	if err != nil {
		return nil, response.ErrCodeFailed
	}
	return &user, response.ErrCodeSuccess
}

// 校验用户Token
func (u *UserService) VerifyToken(token string) error {
	return global.Rdb.Exists(ctx, token).Err()
}
