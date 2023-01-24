package service

import (
	"crm/common"
	"crm/global"
	"crm/models"
	"crm/response"
	"strconv"
	"time"
)

const (
	Open  = 1
	Close = 2
)

type CustomerService struct {
}

// 创建客户
func (c *CustomerService) Create(param *models.CustomerCreateParam) int {
	customer := models.Customer{
		Name:     param.Name,
		Source:   param.Source,
		Phone:    param.Phone,
		Email:    param.Email,
		Industry: param.Industry,
		Level:    param.Level,
		Remarks:  param.Remarks,
		Region:   param.Region,
		Address:  param.Address,
		Status:   1,
		Creator:  param.Creator,
		Created:  time.Now().Unix(),
	}
	if err := global.Db.Create(&customer).Error; err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 更新客户
func (c *CustomerService) Update(param *models.CustomerUpdateParam) int {
	customer := models.Customer{
		Id:       param.Id,
		Name:     param.Name,
		Source:   param.Source,
		Phone:    param.Phone,
		Email:    param.Email,
		Industry: param.Industry,
		Level:    param.Level,
		Remarks:  param.Remarks,
		Region:   param.Region,
		Address:  param.Address,
		Status:   param.Status,
		Updated:  time.Now().Unix(),
	}
	db := global.Db.Model(&customer).Select("*").Omit("id", "creator", "created")
	if err := db.Updates(&customer).Error; err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 发送邮件给客户
func (c *CustomerService) SendMail(param *models.CustomerSendMailParam) int {
	var mc models.MailConfig
	err := global.Db.Model(&models.MailConfig{}).Where("creator = ?", param.Uid).First(&mc).Error
	if err != nil {
		return response.ErrCodeMailSendFailed
	}
	if mc.Status == Close {
		return response.ErrCodeMailSendUnEnable
	}
	mailParam := models.MailParam{
		Smtp:     mc.Stmp,
		Port:     mc.Port,
		AuthCode: mc.AuthCode,
		Sender:   mc.Email,
		Subject:  param.Subject,
		Content:  param.Content,
		Receiver: param.Receiver,
	}
	if err := common.SendMailToCustomer(mailParam); err != nil {
		return response.ErrCodeMailSendFailed
	}
	return response.ErrCodeSuccess
}

// 删除客户
func (c *CustomerService) Delete(param *models.CustomerDeleteParam) int {
	if err := global.Db.Delete(&models.Customer{}, param.Ids).Error; err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 查询客户列表
func (c *CustomerService) QueryList(param *models.CustomerQueryParam) ([]*models.CustomerList, int64, int) {
	customer := models.Customer{
		Name:    param.Name,
		Creator: param.Creator,
	}
	customerList := make([]*models.CustomerList, 0)
	rows, err := restPage(param.Page, CUSTOMER, customer, &customerList, &[]*models.CustomerList{})
	if err != nil {
		return nil, 0, response.ErrCodeFailed
	}
	return customerList, rows, response.ErrCodeSuccess
}

// 查询客户信息
func (c *CustomerService) QueryInfo(param *models.CustomerQueryParam) (*models.CustomerInfo, int) {
	customer := models.Customer{
		Id: param.Id,
	}
	customerInfo := models.CustomerInfo{}
	err := global.Db.Table(CUSTOMER).Where(&customer).First(&customerInfo).Error
	if err != nil {
		return nil, response.ErrCodeFailed
	}
	return &customerInfo, response.ErrCodeSuccess
}

// 查询客户选项
func (c *CustomerService) QueryOption(uid int64) ([]*models.CustomerOption, int) {
	customer := models.Customer{
		Creator: uid,
	}
	option := make([]*models.CustomerOption, 0)
	err := global.Db.Table(CUSTOMER).Where(&customer).Find(&option).Error
	if err != nil {
		return nil, response.ErrCodeFailed
	}
	return option, response.ErrCodeSuccess
}

// 导出Excel文件
func (c *CustomerService) Export(uid int64) (string, int) {
	customers := make([]models.Customer, 0)
	err := global.Db.Where("creator = ?", uid).Find(&customers).Error
	if err != nil {
		return StringNull, response.ErrCodeFailed
	}
	excelRows := make([]models.CustomerExcelRow, 0)
	var row models.CustomerExcelRow
	for _, c := range customers {
		row.Name = c.Name
		row.Source = c.Source
		row.Phone = c.Phone
		row.Email = c.Email
		row.Industry = c.Industry
		row.Level = c.Level
		row.Remarks = c.Remarks
		row.Region = c.Region
		row.Address = c.Address
		if c.Status == 1 {
			row.Status = "已签约"
		}
		if c.Status == 2 {
			row.Status = "未签约"
		}
		row.Created = time.Unix(c.Created, 0).Format("2006-01-02")
		if c.Updated != 0 {
			row.Updated = time.Unix(c.Updated, 0).Format("2006-01-02")
		}
		excelRows = append(excelRows, row)
	}
	sheet := "客户信息"
	columns := []string{"名称", "客户来源", "手机号", "邮箱", "客户行业", "客户级别", "备注", "地区", "详细地址", "成交状态", "创建时间", "更新时间"}
	fileName := "customer_" + strconv.FormatInt(uid, 10)
	file, err := common.GenExcelFile(sheet, columns, excelRows, fileName)
	if err != nil {
		return StringNull, response.ErrCodeFailed
	}
	return file, response.ErrCodeSuccess
}
