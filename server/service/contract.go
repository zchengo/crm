package service

import (
	"crm/global"
	"crm/models"
	"crm/response"

	"time"
)

type ContractService struct {
}

// 创建合同
func (c *ContractService) Create(param *models.ContractCreateParam) int {
	contract := models.Contract{
		Name:        param.Name,
		Amount:      param.Amount,
		BeginTime:   param.BeginTime,
		OverTime:    param.OverTime,
		Remarks:     param.Remarks,
		Cid:         param.Cid,
		Productlist: param.Productlist,
		Status:      param.Status,
		Creator:     param.Creator,
		Created:     time.Now().Unix(),
	}
	if err := global.Db.Create(&contract).Error; err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 更新合同
func (c *ContractService) Update(param *models.ContractUpdateParam) int {
	contract := models.Contract{
		Id:          param.Id,
		Name:        param.Name,
		Amount:      param.Amount,
		BeginTime:   param.BeginTime,
		OverTime:    param.OverTime,
		Remarks:     param.Remarks,
		Cid:         param.Cid,
		Productlist: param.Productlist,
		Status:      param.Status,
		Updated:     time.Now().Unix(),
	}
	db := global.Db.Model(&contract).Select("*").Omit("id", "creator", "created")
	if err := db.Updates(&contract).Error; err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 删除合同
func (c *ContractService) Delete(param *models.ContractDeleteParam) int {
	if err := global.Db.Delete(&models.Contract{}, param.Ids).Error; err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 查询合同列表
func (c *ContractService) QueryList(param *models.ContractQueryParam) ([]*models.ContractList, int64, int) {
	contractList := make([]*models.ContractList, 0)
	s := "contract.id, contract.name, contract.amount, contract.begin_time, contract.over_time, customer.name as cname, contract.remarks, contract.status, contract.created, contract.updated"
	j := "inner join customer on contract.cid = customer.id and contract.creator = ?"

	// 分页查询
	offset := (param.Page.PageNum - 1) * param.Page.PageSize
	mdb := global.Db.Offset(offset).Limit(param.Page.PageSize).Table(CONTRACT).Select(s)
	var err error
	if param.Id != 0 {
		err = mdb.Joins(j+" and contract.id = ?", param.Creator, param.Id).Scan(&contractList).Error
	} else {
		err = mdb.Joins(j, param.Creator).Scan(&contractList).Error
	}
	if err != nil {
		return nil, 0, response.ErrCodeFailed
	}
	var rows int64
	global.Db.Raw("select count(*) from contract where creator = ?", param.Creator).Scan(&rows)
	return contractList, rows, response.ErrCodeSuccess
}

// 查询合同信息
func (c *ContractService) QueryInfo(param *models.ContractQueryParam) (*models.ContractInfo, int) {
	contract := models.Contract{
		Id: param.Id,
	}
	contractInfo := models.ContractInfo{}
	err := global.Db.Table(CONTRACT).Where(&contract).First(&contractInfo).Error
	if err != nil {
		return nil, response.ErrCodeFailed
	}
	return &contractInfo, response.ErrCodeSuccess
}

// 在编辑合同中，添加产品后，返回已添加的产品列表
func (p *ContractService) QueryPlist(param *models.ContractQueryParam) ([]*models.Products, int) {
	if param.Id == 0 {
		products := make([]*models.Products, 0)
		if err := global.Db.Table(PRODUCT).Find(&products, param.Pids).Error; err != nil {
			return nil, response.ErrCodeFailed
		}
		return products, response.ErrCodeSuccess
	}

	// 默认已添加的产品列表
	var contract models.Contract
	err := global.Db.Table(CONTRACT).Select("productlist").First(&contract, param.Id).Error
	if err != nil {
		return nil, response.ErrCodeFailed
	}

	// 最终已添加的产品列表
	addedProductList := make([]*models.Products, 0)

	if len(param.Pids) == 0 {
		return addedProductList, response.ErrCodeSuccess
	}

	addedPids := make([]int64, 0)
	for _, pid := range param.Pids {
		if len(*contract.Productlist) == 0 {
			addedPids = param.Pids
			break
		}
		for _, product := range *contract.Productlist {
			if pid == product.Id {
				addedProductList = append(addedProductList, product)
				continue
			}
			addedPids = append(addedPids, pid)
		}
	}
	products := make([]*models.Products, 0)
	if err := global.Db.Table(PRODUCT).Find(&products, addedPids).Error; err != nil {
		return nil, response.ErrCodeFailed
	}
	addedProductList = append(addedProductList, products...)
	return addedProductList, response.ErrCodeSuccess
}
