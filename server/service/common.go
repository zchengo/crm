package service

import (
	"crm/dao"
	"crm/response"
)

const (

	// 数据库表名
	USER      = "user"
	CUSTOMER  = "customer"
	CONTRACT  = "contract"
	PRODUCT   = "product"
	SUBSCRIBE = "subscribe"
	NOTICE    = "notice"
)

const (
	NumberNull = 0
	StringNull = ""
)

const (
	Dev  = "dev"
	Prod = "prod"
)

const (
	REGISTER_NOTICE_TEMPLATE   = "你注册了账号"
	LOGIN_NOTICE_TEMPLATE      = "你登录了账号"
	SUBSCRIBE_NOTICE_TEMPLATE1 = "你订阅了专业版"
	SUBSCRIBE_NOTICE_TEMPLATE2 = "你订阅了高级版"
)

type CommonService struct {
	commonDao *dao.CommonDao
}

func NewCommonService() *CommonService {
	return &CommonService{
		commonDao: dao.NewCommonDao(),
	}
}

// 初始化数据库
func (c *CommonService) InitDatabase() int {
	if err := c.commonDao.InitDatabase(); err != nil {
		return response.ErrCodeInitDataFailed
	}
	return response.ErrCodeSuccess
}
