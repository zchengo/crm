package service

import (
	"context"
	"crm/global"
	"crm/models"
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
)

const (
	NumberNull = 0
	StringNull = ""
)

const (
	Dev  = "dev"
	Prod = "prod"
)

var ctx = context.Background()

// RestPage 分页查询
// page  设置起始页、每页条数,
// name  查询目标表的名称
// query 查询条件,
// dest  查询结果绑定的结构体,
// bind  绑定表结构对应的结构体
func restPage(page models.Page, name string, query interface{}, dest interface{}, bind interface{}) (int64, error) {
	if page.PageNum > 0 && page.PageSize > 0 {
		offset := (page.PageNum - 1) * page.PageSize
		global.Db.Offset(offset).Limit(page.PageSize).Table(name).Where(query).Find(dest)
	}
	res := global.Db.Table(name).Where(query).Find(bind)
	return res.RowsAffected, res.Error
}

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
