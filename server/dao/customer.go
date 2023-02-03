package dao

import (
	"crm/global"
	"crm/models"
	"time"
)

type CustomerDao struct {
}

func NewCustomerDao() *CustomerDao {
	return &CustomerDao{}
}

func (c *CustomerDao) Create(param *models.CustomerCreateParam) error {
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
	return global.Db.Create(&customer).Error
}

func (c *CustomerDao) Update(param *models.CustomerUpdateParam) error {
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
	return db.Updates(&customer).Error
}

func (c *CustomerDao) Delete(param *models.CustomerDeleteParam) error {
	return global.Db.Delete(&models.Customer{}, param.Ids).Error
}

func (c *CustomerDao) IsExists(name string, uid int64) bool {
	var customer models.Customer
	db := global.Db.Table(CUSTOMER).Where("name = ? and creator = ?", name, uid).First(&customer)
	return db.RowsAffected != NumberNull
}

func (c *CustomerDao) GetList(param *models.CustomerQueryParam) ([]*models.CustomerList, int64, error) {
	customer := models.Customer{
		Name:     param.Name,
		Source:   param.Source,
		Industry: param.Industry,
		Level:    param.Level,
		Status:   param.Status,
		Creator:  param.Creator,
	}
	customerList := make([]*models.CustomerList, 0)
	rows, err := restPage(param.Page, CUSTOMER, customer, &customerList, &[]*models.CustomerList{})
	if err != nil {
		return nil, 0, err
	}
	return customerList, rows, nil
}

func (c *CustomerDao) GetInfo(param *models.CustomerQueryParam) (*models.CustomerInfo, error) {
	customer := models.Customer{
		Id: param.Id,
	}
	customerInfo := models.CustomerInfo{}
	err := global.Db.Table(CUSTOMER).Where(&customer).First(&customerInfo).Error
	if err != nil {
		return nil, err
	}
	return &customerInfo, nil
}

func (c *CustomerDao) GetOption(uid int64) ([]*models.CustomerOption, error) {
	customer := models.Customer{
		Creator: uid,
	}
	option := make([]*models.CustomerOption, 0)
	err := global.Db.Table(CUSTOMER).Where(&customer).Find(&option).Error
	if err != nil {
		return nil, err
	}
	return option, nil
}

func (c *CustomerDao) GetListByUid(uid int64) ([]*models.Customer, error) {
	customers := make([]*models.Customer, 0)
	err := global.Db.Where("creator = ?", uid).Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}
