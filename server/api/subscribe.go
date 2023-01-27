package api

import (
	"crm/common"
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
	err := context.ShouldBind(&param)
	if int64(uid) <= 0 || err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	param.Uid = int64(uid)
	payUrl, errCode := s.subscribeService.Pay(param)
	response.Result(errCode, payUrl, context)
}

// 支付成功回调
func (s *SubscribeApi) PayBack(context *gin.Context) {
	notifyReq := common.GetAlipay().VerifySign(context.Request)
	errCode := s.subscribeService.PayBack(notifyReq.GetString("out_trade_no"))
	context.String(http.StatusOK, "%s", "success")
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
