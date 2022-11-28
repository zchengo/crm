package service

import (
	"crm/global"
	"crm/models"
	"fmt"
	"time"
)

type DashboardService struct {
}

// 数据汇总
func (d *DashboardService) Summary(uid int64) models.DashboardSum {
	var ds models.DashboardSum
	global.Db.Raw("select count(*) from customer where creator = ?", uid).Scan(&ds.Customers)
	global.Db.Raw("select count(*) from contract where creator = ?", uid).Scan(&ds.Contracts)
	global.Db.Raw("select sum(amount) from contract where creator = ?", uid).Scan(&ds.ContractAmount)
	global.Db.Raw("select count(*) from product where creator = ?", uid).Scan(&ds.Products)

	now := time.Now().Unix()
	con := "created > ? and created < ? and status = 1 and creator = ?"
	for i, d := 0, 7; i < 7; i++ {
		day := now - (int64(d) * 24 * 60 * 60)
		start, end := dayRange(day)
		fmt.Println(start, end)
		global.Db.Table(CONTRACT).Where(con, start, end, uid).Pluck("amount", &ds.Amount[i])
		ds.Date[i] = time.Unix(day, 0).Format("01-02")
		d--
	}
	return ds
}

// 获取某一天的时间范围
func dayRange(day int64) (int64, int64) {
	t := time.Unix(day, 0)
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	end := start.AddDate(0, 0, 1)
	return start.Unix(), end.Unix()
}
