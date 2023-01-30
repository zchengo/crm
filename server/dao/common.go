package dao

import (
	"context"
	"crm/common"
	"crm/global"
	"crm/models"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

const (

	// 数据库表名
	USER        = "user"
	CUSTOMER    = "customer"
	CONTRACT    = "contract"
	PRODUCT     = "product"
	SUBSCRIBE   = "subscribe"
	NOTICE      = "notice"
	MAIL_CONFIG = "mail_config"

	// 空值
	NumberNull = 0
	StringNull = ""

	// 运行环境
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

type CommonDao struct {
}

func NewCommonDao() *CommonDao {
	return &CommonDao{}
}

func (c *CommonDao) InitDatabase() error {
	env := global.Config.Server.Runenv
	if env == Prod {
		dbFile := global.Config.Mysql.DbFile
		sql, err := os.ReadFile(dbFile)
		if err != nil {
			log.Printf("Common.InitDatabase.Error: read file %s error: %s", dbFile, err)
			return err
		}
		sqls := strings.Split(string(sql), ";")
		for _, sql := range sqls {
			s := strings.TrimSpace(sql)
			if s == StringNull {
				continue
			}
			if err := global.Db.Exec(s).Error; err != nil {
				log.Printf("Common.InitDatabase.Error: %s", err)
				return err
			}
		}
		return nil
	}
	return nil
}

func (c *CommonDao) FileUpload(file *multipart.FileHeader) (*models.FileInfo, error) {
	dist := global.Config.File.Path
	name := common.GenUUID() + path.Ext(file.Filename)
	dn := dist + name
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	out, err := os.Create(dn)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return nil, err
	}
	flieInfo := models.FileInfo{
		Url:  dn,
		Name: name,
	}
	return &flieInfo, err
}

func (c *CommonDao) FileRemove(fileName string) error {
	file := global.Config.File.Path + fileName
	return os.Remove(file)
}
