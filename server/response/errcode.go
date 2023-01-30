package response

const (
	ErrCodeSuccess      = 0 // 成功
	ErrCodeFailed       = 1 // 失败
	ErrCodeParamInvalid = 2 // 请求参数无效
	ErrCodeNoLogin      = 3 // 未登录或非法访问
	ErrCodeTokenExpire  = 4 // Token过期

	ErrCodeInitDataFailed   = 10 // 初始化数据失败
	ErrCodeFileUploadFailed = 11 // 文件上传失败
	ErrCodeFileRemoveFailed = 12 // 文件移除失败

	ErrCodeUserHasExist         = 10001 // 用户已经存在
	ErrCodeUserNotExist         = 10002 // 用户不存在
	ErrCOdeUserEmailOrPass      = 10003 // 用户邮箱或密码错误
	ErrCodeVerityCodeSendFailed = 10004 // 验证码发送失败
	ErrCodeVerityCodeInvalid    = 10005 // 验证码无效
	ErrCodeEmailFormatInvalid   = 10008 // 邮箱格式无效
	ErrCodeUserPassResetFailed  = 10009 // 用户密码重置失败

	ErrCodeCustomerHasExist = 20001 // 客户名称已经存在

	ErrCodePayFailed = 20001 // 支付宝支付失败

	ErrCodeFileExportFailed = 30001 // 文件导出失败

	ErrCodeProductHasExist = 40001 // 产品名称已经存在

	ErrCodeMailConfigInvalid = 50001 // 邮件服务配置无效
	ErrCodeMailSendFailed    = 50002 // 邮件发送失败
	ErrCodeMailSendUnEnable  = 50003 // 邮件服务未开启
	ErrCodeMailConfigUnSet   = 50004 // 邮件服务未配置
)

var msg = map[int]string{
	ErrCodeSuccess:              "success",
	ErrCodeFailed:               "failed",
	ErrCodeParamInvalid:         "param invalid",
	ErrCodeNoLogin:              "no login",
	ErrCodeTokenExpire:          "token expire",
	ErrCodeInitDataFailed:       "data init failed",
	ErrCodeFileUploadFailed:     "file upload failed",
	ErrCodeUserHasExist:         "user has exist",
	ErrCodeUserNotExist:         "user not exist",
	ErrCOdeUserEmailOrPass:      "user email or password error",
	ErrCodeVerityCodeSendFailed: "verify code send failed",
	ErrCodeVerityCodeInvalid:    "verify code invalid",
	ErrCodeEmailFormatInvalid:   "email format invalid",
	ErrCodeUserPassResetFailed:  "user password reset failed",
	ErrCodeFileExportFailed:     "file export failed",
	ErrCodeProductHasExist:      "product has exist",
	ErrCodeMailConfigInvalid:    "mail config invalid",
	ErrCodeMailSendFailed:       "mail send failed",
	ErrCodeMailSendUnEnable:     "mail send server unEnable",
	ErrCodeMailConfigUnSet:      "mail config un set",
}
