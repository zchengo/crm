package api

import (
	"crm/response"
	"crm/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DashboardApi struct {
	dashboardService *service.DashboardService
}

func NewDashboardApi() *DashboardApi {
	dashboardApi := DashboardApi{
		dashboardService: &service.DashboardService{},
	}
	return &dashboardApi
}

// 获取数据汇总
func (d *DashboardApi) Summary(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	sum := d.dashboardService.Summary(int64(uid))
	response.Result(response.ErrCodeSuccess, sum, context)
}
