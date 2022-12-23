package service

import (
	"crm/global"
	"crm/models"
	"crm/response"
	"time"
)

type ProductService struct {
}

// 创建产品
func (p *ProductService) Create(param *models.ProductCreateParam) int {
	product := models.Product{
		Name:        param.Name,
		Type:        param.Type,
		Unit:        param.Unit,
		Code:        param.Code,
		Price:       param.Price,
		Description: param.Description,
		Status:      param.Status,
		Creator:     param.Creator,
		Created:     time.Now().Unix(),
	}
	if err := global.Db.Create(&product).Error; err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 更新产品
func (p *ProductService) Update(param *models.ProductUpdateParam) int {
	product := models.Product{
		Id:          param.Id,
		Name:        param.Name,
		Type:        param.Type,
		Unit:        param.Unit,
		Code:        param.Code,
		Price:       param.Price,
		Description: param.Description,
		Status:      param.Status,
		Updated:     time.Now().Unix(),
	}
	db := global.Db.Model(&product).Select("*").Omit("id", "creator", "created")
	if err := db.Updates(&product).Error; err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 删除产品
func (p *ProductService) Delete(param *models.ProductDeleteParam) int {
	err := global.Db.Delete(&models.Product{}, param.Ids).Error
	if err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 查询产品列表
func (p *ProductService) QueryList(param *models.ProductQueryParam) ([]*models.ProductList, int64, int) {
	product := models.Product{
		Name:    param.Name,
		Status:  param.Status,
		Creator: param.Creator,
	}
	productList := make([]*models.ProductList, 0)
	rows, err := restPage(param.Page, PRODUCT, product, &productList, &[]*models.ProductList{})
	if err != nil {
		return nil, 0, response.ErrCodeFailed
	}
	return productList, rows, response.ErrCodeSuccess
}

// 查询产品信息
func (p *ProductService) QueryInfo(param *models.ProductQueryParam) (*models.ProductInfo, int) {
	product := models.Product{
		Id: param.Id,
	}
	productInfo := models.ProductInfo{}
	err := global.Db.Table(PRODUCT).Where(&product).First(&productInfo).Error
	if err != nil {
		return nil, response.ErrCodeFailed
	}
	return &productInfo, response.ErrCodeSuccess
}
