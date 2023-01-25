package dao

import (
	"crm/global"
	"crm/models"
)

type DashboardDao struct {
}

func NewDashboardDao() *DashboardDao {
	return &DashboardDao{}
}

func (d *DashboardDao) Count(uid int64, days int) models.DashboardSum {
	var ds models.DashboardSum
	global.Db.Raw("select count(*) from customer where creator = ?", uid).Scan(&ds.Customers)
	global.Db.Raw("select count(*) from contract where creator = ?", uid).Scan(&ds.Contracts)
	global.Db.Raw("select sum(amount) from contract where creator = ?", uid).Scan(&ds.ContractAmount)
	global.Db.Raw("select count(*) from product where creator = ?", uid).Scan(&ds.Products)

	s := "industry as name, count(*) as value"
	global.Db.Model(&models.Customer{}).Select(s).Where("creator = ?", uid).Group("industry").Find(&ds.CustomerIndustry)
	return ds
}

func (d *DashboardDao) AmountSum(start, end, uid int64) float64 {
	var as float64
	con := "created > ? and created < ? and status = 1 and creator = ?"
	global.Db.Table(CONTRACT).Where(con, start, end, uid).Pluck("amount", &as)
	return as
}
