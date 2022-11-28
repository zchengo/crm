package api

import (
	"crm/models"
	"crm/response"
	"crm/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerApi struct {
	customerService *service.CustomerService
}

func NewCustomerApi() *CustomerApi {
	customerApi := CustomerApi{
		customerService: &service.CustomerService{},
	}
	return &customerApi
}

// 创建产品
func (c *CustomerApi) Create(context *gin.Context) {
	var param models.CustomerCreateParam
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	err := context.ShouldBind(&param)
	if uid <= 0 || err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	param.Creator = int64(uid)
	errCode := c.customerService.Create(&param)
	response.Result(errCode, nil, context)
}

// 更新产品
func (c *CustomerApi) Update(context *gin.Context) {
	var param models.CustomerUpdateParam
	if err := context.ShouldBind(&param); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	errCode := c.customerService.Update(&param)
	response.Result(errCode, nil, context)
}

// 删除客户
func (c *CustomerApi) Delete(context *gin.Context) {
	var param models.CustomerDeleteParam
	if err := context.ShouldBind(&param); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	errCode := c.customerService.Delete(&param)
	response.Result(errCode, nil, context)
}

// 查询客户列表
func (c *CustomerApi) QueryList(context *gin.Context) {
	var param models.CustomerQueryParam
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	err := context.ShouldBind(&param)
	if uid <= 0 || err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	param.Creator = int64(uid)
	customerList, rows, errCode := c.customerService.QueryList(&param)
	response.PageResult(errCode, customerList, rows, context)
}

// 查询客户信息
func (c *CustomerApi) QueryInfo(context *gin.Context) {
	var param models.CustomerQueryParam
	if err := context.ShouldBind(&param); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	customerInfo, errCode := c.customerService.QueryInfo(&param)
	response.Result(errCode, customerInfo, context)
}

// 查询客户选项
func (c *CustomerApi) QueryOption(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	if uid <= 0 {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	customerOption, errCode := c.customerService.QueryOption(int64(uid))
	response.Result(errCode, customerOption, context)
}
