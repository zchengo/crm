package api

import (
	"crm/response"
	"crm/service"

	"github.com/gin-gonic/gin"
)

type CommonApi struct {
	commonService *service.CommonService
}

func NewCommonApi() *CommonApi {
	return &CommonApi{
		commonService: service.NewCommonService(),
	}
}

// 初始化数据库
func (c *CommonApi) InitDatabase(context *gin.Context) {
	errCode := c.commonService.InitDatabase()
	response.Result(errCode, nil, context)
}