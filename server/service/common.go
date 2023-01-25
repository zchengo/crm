package service

import (
	"crm/global"
	"crm/response"
	"os"

	"log"
	"strings"
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

// 初始化数据库数据
func InitData() int {
	env := global.Config.Server.Runenv
	if env == Prod {
		fn := "/home/ubuntu/crmapi/crm.sql"
		sql, err := os.ReadFile(fn)
		if err != nil {
			log.Printf("[ERROR] read file %s error: %s", fn, err)
			return response.ErrCodeInitDataFailed
		}
		sqls := strings.Split(string(sql), ";")
		for _, sql := range sqls {
			s := strings.TrimSpace(sql)
			if s == StringNull {
				continue
			}
			if err := global.Db.Exec(s).Error; err != nil {
				log.Printf("[ERROR] datebase flie import error: %s", err)
				return response.ErrCodeInitDataFailed
			}
		}
		return response.ErrCodeInitDataSuccess
	}
	return response.ErrCodeSuccess
}
