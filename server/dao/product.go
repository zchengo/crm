package dao

import (
	"crm/global"
	"crm/models"
	"time"
)

type ProductDao struct {
}

func NewProductDao() *ProductDao {
	return &ProductDao{}
}

func (p *ProductDao) Create(param *models.ProductCreateParam) error {
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
	return global.Db.Create(&product).Error
}

func (p *ProductDao) Update(param *models.ProductUpdateParam) error {
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
	return db.Updates(&product).Error
}

func (p *ProductDao) Delete(param *models.ProductDeleteParam) error {
	return global.Db.Delete(&models.Product{}, param.Ids).Error
}

func (p *ProductDao) IsExists(name string, uid int64) bool {
	var product models.Product
	db := global.Db.Table(PRODUCT).Where("name = ? and creator = ?", name, uid).First(&product)
	return db.RowsAffected != NumberNull
}

func (p *ProductDao) GetList(param *models.ProductQueryParam) ([]*models.ProductList, int64, error) {
	product := models.Product{
		Name:    param.Name,
		Status:  param.Status,
		Creator: param.Creator,
	}
	productList := make([]*models.ProductList, 0)
	rows, err := restPage(param.Page, PRODUCT, product, &productList, &[]*models.ProductList{})
	if err != nil {
		return nil, 0, err
	}
	return productList, rows, err
}

func (p *ProductDao) GetListByIds(ids []int64) ([]*models.Products, error) {
	products := make([]*models.Products, 0)
	if err := global.Db.Table(PRODUCT).Find(&products, ids).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductDao) GetInfo(param *models.ProductQueryParam) (*models.ProductInfo, error) {
	product := models.Product{
		Id: param.Id,
	}
	productInfo := models.ProductInfo{}
	err := global.Db.Table(PRODUCT).Where(&product).First(&productInfo).Error
	if err != nil {
		return nil, err
	}
	return &productInfo, err
}

func (p *ProductDao) GetListByUid(uid int64) ([]*models.Product, error) {
	products := make([]*models.Product, 0)
	err := global.Db.Where("creator = ?", uid).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
