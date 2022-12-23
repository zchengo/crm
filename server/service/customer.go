package service

import (
	"crm/global"
	"crm/models"
	"crm/response"
	"time"
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
