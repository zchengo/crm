package dao

import (
	"crm/global"
	"crm/models"
	"time"
)

type ContractDao struct {
}

func NewContractDao() *ContractDao {
	return &ContractDao{}
}

func (c *ContractDao) Create(param *models.ContractCreateParam) error {
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
	return global.Db.Create(&contract).Error
}

func (c *ContractDao) Update(param *models.ContractUpdateParam) error {
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
	return db.Updates(&contract).Error
}

func (c *ContractDao) Delete(param *models.ContractDeleteParam) error {
	return global.Db.Delete(&models.Contract{}, param.Ids).Error
}

func (c *ContractDao) GetList(param *models.ContractQueryParam) ([]*models.ContractList, int64, error) {
	contractList := make([]*models.ContractList, 0)
	field := "contract.id, contract.name, contract.amount, contract.begin_time, contract.over_time, customer.name as cname, contract.remarks, contract.status, contract.created, contract.updated"
	where := "inner join customer on contract.cid = customer.id and contract.creator = ?"
	raw := "select count(*) from contract where creator = ?"

	// 分页查询
	offset := (param.Page.PageNum - 1) * param.Page.PageSize
	db := global.Db.Offset(offset).Limit(param.Page.PageSize).Table(CONTRACT).Select(field)
	
	var rows int64
	if param.Id != NumberNull {
		db.Joins(where+" and contract.id = ?", param.Creator, param.Id)
		global.Db.Raw(raw + " and contract.id = ?", param.Creator, param.Creator).Scan(&rows)
	} else {
		if param.Status != NumberNull {
			db.Joins(where+" and contract.status = ?", param.Creator, param.Status)
			global.Db.Raw(raw + " and contract.status = ?", param.Creator, param.Status).Scan(&rows)
		} else {
			db.Joins(where, param.Creator)
			global.Db.Raw(raw, param.Creator).Scan(&rows)
		}
	}
	if err := db.Scan(&contractList).Error; err != nil {
		return nil, NumberNull, nil
	}
	return contractList, rows, nil
}

func (c *ContractDao) GetListByUid(uid int64) ([]*models.ContractList, error) {
	contracts := make([]*models.ContractList, 0)
	s := "contract.id, contract.name, contract.amount, contract.begin_time, contract.over_time, customer.name as cname, contract.remarks, contract.status, contract.created, contract.updated"
	j := "left join customer on contract.cid = customer.id and contract.creator = ?"
	err := global.Db.Table(CONTRACT).Select(s).Joins(j, uid).Scan(&contracts).Error
	if err != nil {
		return nil, err
	}
	return contracts, nil
}

func (c *ContractDao) GetInfo(param *models.ContractQueryParam) (*models.ContractInfo, error) {
	contract := models.Contract{
		Id: param.Id,
	}
	contractInfo := models.ContractInfo{}
	err := global.Db.Table(CONTRACT).Where(&contract).First(&contractInfo).Error
	if err != nil {
		return nil, err
	}
	return &contractInfo, nil
}

func (c *ContractDao) GetAddedPList(id int64) (*models.Contract, error) {
	var contract models.Contract
	err := global.Db.Table(CONTRACT).Select("productlist").First(&contract, id).Error
	if err != nil {
		return nil, err
	}
	return &contract, nil
}
