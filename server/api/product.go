package api

import (
	"crm/models"
	"crm/response"
	"crm/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductApi struct {
	productService *service.ProductService
}

func NewProductApi() *ProductApi {
	productApi := ProductApi{
		productService: &service.ProductService{},
	}
	return &productApi
}

// 创建产品
func (p *ProductApi) Create(context *gin.Context) {
	var param models.ProductCreateParam
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	err := context.ShouldBind(&param)
	if uid <= 0 || err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	param.Creator = int64(uid)
	errCode := p.productService.Create(&param)
	response.Result(errCode, nil, context)
}

// 更新产品
func (p *ProductApi) Update(context *gin.Context) {
	var param models.ProductUpdateParam
	if err := context.ShouldBind(&param); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	errCode := p.productService.Update(&param)
	response.Result(errCode, nil, context)
}

// 删除产品
func (p *ProductApi) Delete(context *gin.Context) {
	var param models.ProductDeleteParam
	if err := context.ShouldBind(&param); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	errCode := p.productService.Delete(&param)
	response.Result(errCode, nil, context)
}

// 查询产品列表
func (p *ProductApi) GetList(context *gin.Context) {
	var param models.ProductQueryParam
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	err := context.ShouldBind(&param)
	if uid <= 0 || err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	param.Creator = int64(uid)
	productList, rows, errCode := p.productService.GetList(&param)
	response.PageResult(errCode, productList, rows, context)
}

// 查询产品信息
func (p *ProductApi) GetInfo(context *gin.Context) {
	var param models.ProductQueryParam
	if err := context.ShouldBind(&param); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	productInfo, errCode := p.productService.GetInfo(&param)
	response.Result(errCode, productInfo, context)
}

// 导出Excel文件
func (p *ProductApi) Export(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	if uid <= 0 {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	file, errCode := p.productService.Export(int64(uid))
	if len(file) >= 0 && errCode != 0 {
		response.Result(errCode, nil, context)
		return
	}
	context.File(file)
}
