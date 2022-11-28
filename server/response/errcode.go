package response

const (
	ErrCodeSuccess      = 0 // 成功
	ErrCodeFailed       = 1 // 失败
	ErrCodeParamInvalid = 2 // 请求参数无效
	ErrCodeNoLogin      = 3 // 未登录或非法访问
	ErrCodeTokenExpire  = 4 // Token过期

	ErrCodeUserHasExist         = 10001 // 用户已经存在
	ErrCodeUserNotExist         = 10002 // 用户不存在
	ErrCOdeUserEmailOrPass      = 10003 // 用户邮箱或密码错误
	ErrCodeVerityCodeSendFailed = 10004 // 验证码发送失败
	ErrCodeVerityCodeInvalid    = 10005 // 验证码无效
	ErrCodeCompanyCreateFailed  = 10006 // 企业创建失败
	ErrCodeCompanyIdNotExist    = 10007 // 企业编号不存在
	ErrCodeEmailFormatInvalid   = 10008 // 邮箱格式无效
	ErrCodeUserPassResetFailed  = 10009 // 用户密码重置失败
)

var msg = map[int]string{
	ErrCodeSuccess:              "success",
	ErrCodeFailed:               "failed",
	ErrCodeParamInvalid:         "param invalid",
	ErrCodeNoLogin:              "no login",
	ErrCodeTokenExpire:          "token expire",
	ErrCodeUserHasExist:         "user has exist",
	ErrCodeUserNotExist:         "user not exist",
	ErrCOdeUserEmailOrPass:      "user email or password error",
	ErrCodeVerityCodeSendFailed: "verify code send failed",
	ErrCodeVerityCodeInvalid:    "verify code invalid",
	ErrCodeCompanyCreateFailed:  "company create failed",
	ErrCodeEmailFormatInvalid:   "email format invalid",
	ErrCodeUserPassResetFailed:  "user password reset failed",
}
