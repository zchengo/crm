package api

import (
	"crm/models"
	"crm/response"
	"crm/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubscribeApi struct {
	subscribeService *service.SubscribeService
}

func NewSubscribeApi() *SubscribeApi {
	subscribeApi := SubscribeApi{
		subscribeService: service.NewSubscribeService(),
	}
	return &subscribeApi
}

// 订阅专业版，发起支付
func (s *SubscribeApi) Pay(context *gin.Context) {
	var param models.SubscribePayParam
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	err := context.ShouldBind(&param);
	if int64(uid) <= 0 ||  err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	param.Uid = int64(uid)
	payUrl, errCode := s.subscribeService.Pay(param)
	response.Result(errCode, payUrl, context)
}

// 支付成功回调
func (s *SubscribeApi) Callback(context *gin.Context) {
	context.Request.ParseForm()
	var outTradeNo = context.Request.Form.Get("out_trade_no")
	paySuccessURL, _ := s.subscribeService.Callback(outTradeNo)
	context.Redirect(http.StatusMovedPermanently, paySuccessURL)
}

// 支付通知
func (s *SubscribeApi) Notify(context *gin.Context) {
	context.Request.ParseForm()
	var outTradeNo = context.Request.Form.Get("out_trade_no")
	errCode := s.subscribeService.Notify(context.Request.Form, outTradeNo)
	response.Result(errCode, nil, context)
}

// 获取订阅信息
func (s *SubscribeApi) GetInfo(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	if int64(uid) <= 0 {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	subscribeInfo, errCode := s.subscribeService.GetInfo(int64(uid))
	response.Result(errCode, subscribeInfo, context)
}