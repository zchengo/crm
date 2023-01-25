package service

import (
	"crm/dao"
	"crm/models"
	"time"
)

type DashboardService struct {
	dashboardDao *dao.DashboardDao
}

func NewDashboardService() *DashboardService {
	return &DashboardService{
		dashboardDao: dao.NewDashboardDao(),
	}
}

// 数据汇总
func (d *DashboardService) Summary(uid int64, days int) models.DashboardSum {
	ds := d.dashboardDao.Count(uid, days)
	now := time.Now().Unix()
	ds.Amount = make([]float64, days)
	ds.Date = make([]string, days)
	for i, dd := 0, days; i < days; i++ {
		day := now - (int64(dd) * 24 * 60 * 60)
		start, end := dayRange(day)
		ds.Amount[i] = d.dashboardDao.AmountSum(start, end, uid)
		ds.Date[i] = time.Unix(day, 0).Format("01-02")
		dd--
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
