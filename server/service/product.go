package service

import (
	"crm/common"
	"crm/global"
	"crm/models"
	"crm/response"
	"strconv"
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

// 导出Excel文件
func (p *ProductService) Export(uid int64) (string, int) {
	products := make([]models.Product, 0)
	err := global.Db.Where("creator = ?", uid).Find(&products).Error
	if err != nil {
		return StringNull, response.ErrCodeFileExportFailed
	}
	excelRows := make([]models.ProductExcelRow, 0)
	var row models.ProductExcelRow
	for _, p := range products {
		row.Name = p.Name
		if p.Status == 1 {
			row.Status = "已上架"
		}
		if p.Status == 2 {
			row.Status = "已下架"
		}
		if p.Type == 1 {
			row.Type = "默认"
		}
		row.Unit = p.Unit
		row.Code = p.Code
		row.Price = p.Price
		row.Description = p.Description
		row.Created = time.Unix(p.Created, 0).Format("2006-01-02")
		if p.Updated != 0 {
			row.Updated = time.Unix(p.Updated, 0).Format("2006-01-02")
		}
		excelRows = append(excelRows, row)
	}
	sheet := "产品信息"
	columns := []string{"名称", "是否上下架", "类型", "单位", "编码", "价格", "描述", "创建时间", "更新时间"}
	fileName := "product_" + strconv.FormatInt(uid, 10)
	file, err := common.GenExcelFile(sheet, columns, excelRows, fileName)
	if err != nil {
		return StringNull, response.ErrCodeFileExportFailed
	}
	return file, response.ErrCodeSuccess
}
