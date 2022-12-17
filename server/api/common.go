package api

import (
	"crm/response"
	"crm/service"

	"github.com/gin-gonic/gin"
)

// 初始化数据
func InitData(c *gin.Context) {
	errCode := service.InitData()
	response.Result(errCode, nil, c)
}