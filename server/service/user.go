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
)

const (
	TOKEN_MAX_EXPIRE_TIME = 7 * 24 // Token最长有效期
)

type UserService struct {
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
		Version:  1,
		Status:   1,
		Created:  time.Now().Unix(),
	}
	if err := global.Db.Create(&newUser).Error; err != nil {
		return response.ErrCodeFailed
	}
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

// 修改邮箱
func (u *UserService) UpdateMail(param *models.UserMailParam) int {
	// 校验验证码是否正确
	key := fmt.Sprintf("user:%s:code", param.Email)
	code := global.Rdb.Get(ctx, key).Val()
	if code != param.Code {
		return response.ErrCodeVerityCodeInvalid
	}
	user := models.User{
		Email:   param.NewEmail,
		Updated: time.Now().Unix(),
	}
	err := global.Db.Table(USER).Where("email = ?", param.Email).Updates(&user).Error
	if err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 退出登录
func (u *UserService) Logout(token string) int {
	err := global.Rdb.Del(ctx, token).Err()
	if err != nil {
		return response.ErrCodeFailed
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
	err := global.Db.Delete(&models.User{}, param.Id).Error
	if err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 获取用户信息
func (u *UserService) GetInfo(uid int64) (*models.UserPersonInfo, int) {
	var user models.UserPersonInfo
	err := global.Db.Table(USER).Where("id = ?", uid).First(&user).Error
	if err != nil {
		return nil, response.ErrCodeFailed
	}

	// 判断用户订阅是否过期
	if user.Version == 2 && time.Now().Unix() > int64(user.Expired) {
		err := global.Db.Model(&models.User{}).Where("id = ?", uid).Update("version", 1).Error
		if err != nil {
			return nil, response.ErrCodeFailed
		}
		return &user, response.ErrCodeSuccess
	}
	return &user, response.ErrCodeSuccess
}

// 订阅个人版
func (u *UserService) Buy(uid int64) (*models.UserVerisonInfo, int) {
	// 订阅一个月按30天计算
	month := time.Now().Unix() + int64(2592000)
	user := models.User{
		Version: 2,
		Expired: month,
		Updated: time.Now().Unix(),
	}
	err := global.Db.Model(&models.User{}).Where("id = ?", uid).Updates(&user).Error
	if err != nil {
		return nil, response.ErrCodeFailed
	}
	// 查询版本信息
	var version models.UserVerisonInfo
	if err := global.Db.Table(USER).Where("id = ?", uid).First(&version).Error; err != nil {
		return nil, response.ErrCodeFailed
	}
	return &version, response.ErrCodeSuccess
}

// 校验用户Token
func (u *UserService) VerifyToken(token string) error {
	return global.Rdb.Exists(ctx, token).Err()
}
