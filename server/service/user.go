package service

import (
	"crm/common"
	"crm/dao"
	"crm/models"
	"crm/response"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

const (
	TOKEN_MAX_EXPIRE_TIME = 7 * 24 // Token最长有效期
)

type UserService struct {
	userDao      *dao.UserDao
	subscribeDao *dao.SubscribeDao
	noticeDao    *dao.NoticeDao
}

func NewUserService() *UserService {
	userService := UserService{
		userDao:      dao.NewUserDao(),
		subscribeDao: dao.NewSubscribeDao(),
		noticeDao:    dao.NewNoticeDao(),
	}
	return &userService
}

// 用户注册
func (u *UserService) Register(param *models.UserCreateParam) int {

	// 判断用户是否存在
	if u.userDao.IsExists(param.Email) {
		return response.ErrCodeUserHasExist
	}

	// 校验验证码是否正确
	code := u.userDao.GetCode(param.Email)
	if code != param.Code {
		return response.ErrCodeVerityCodeInvalid
	}

	// 对密码进行加密
	password, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.ErrCodeFailed
	}
	param.Password = string(password)

	// 创建用户
	if err := u.userDao.Create(param); err != nil {
		return response.ErrCodeFailed
	}
	// 获取UID
	uid, err := u.userDao.GetUid(param.Email)
	if err != nil {
		return response.ErrCodeFailed
	}

	// 新用户默认订阅免费版
	if !u.subscribeDao.IsExists(uid) {
		subscribe := models.SubscribeCreateParam{
			Uid:     uid,
			Version: 1,
		}
		if err := u.subscribeDao.Create(&subscribe); err != nil {
			return response.ErrCodeFailed
		}
	} else {
		subscribe := models.SubscribeUpdateParam{
			Uid:     uid,
			Version: 1,
		}
		if err := u.subscribeDao.Update(&subscribe); err != nil {
			return response.ErrCodeFailed
		}
	}

	// 注册通知
	u.noticeDao.Create(&models.NoticeCreateParam{
		Content: REGISTER_NOTICE_TEMPLATE,
		Creator: uid,
	})

	return response.ErrCodeSuccess
}

// 用户登录
func (u *UserService) Login(param *models.UserLoginParam) (*models.UserInfo, int) {

	// 判断用户是否存在
	if !u.userDao.IsExists(param.Email) {
		return nil, response.ErrCodeUserNotExist
	}

	// 获取用户信息
	user, err := u.userDao.GetUser(param.Email)
	if err != nil {
		return nil, response.ErrCodeFailed
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
	u.noticeDao.Create(&models.NoticeCreateParam{
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
	if err := u.userDao.SetCode(code, email); err != nil {
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
	if !u.userDao.IsExists(param.Email) {
		return response.ErrCodeUserNotExist
	}

	// 校验验证码是否正确
	code := u.userDao.GetCode(param.Email)
	if code != param.Code {
		return response.ErrCodeVerityCodeInvalid
	}

	// 对密码进行加密，更新密码
	password, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.ErrCodeFailed
	}
	if err := u.userDao.UpdatePass(param.Email, string(password)); err != nil {
		return response.ErrCodeUserPassResetFailed
	}
	return response.ErrCodeSuccess
}

// 注销账号
func (u *UserService) Delete(param models.UserDeleteParam) int {
	// 校验验证码是否正确
	code := u.userDao.GetCode(param.Email)
	if code != param.Code {
		return response.ErrCodeVerityCodeInvalid
	}
	// 删除用户所有数据
	if err := u.userDao.DeleteData(param); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 获取用户信息
func (u *UserService) GetInfo(uid int64) (*models.UserPersonInfo, int) {
	userInfo, err := u.userDao.GetInfo(uid)
	if err != nil {
		return nil, response.ErrCodeFailed
	}
	return userInfo, response.ErrCodeSuccess
}
