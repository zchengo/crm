package api

import (
	"crm/models"
	"crm/response"
	"crm/service"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NoticeApi struct {
	noticeService *service.NoticeService
}

func NewNoticeApi() *NoticeApi {
	noticeApi := NoticeApi{
		noticeService: &service.NoticeService{},
	}
	return &noticeApi
}

func (n *NoticeApi) Update(context *gin.Context) {
	var param models.NoticeUpdateParam
	if err := context.ShouldBind(&param); err != nil {
		fmt.Println(err)
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	errCode := n.noticeService.Update(&param)
	response.Result(errCode, nil, context)
}

func (n *NoticeApi) GetUnReadCount(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	if uid <= 0 {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	unReadNotice, errCode := n.noticeService.GetUnReadCount(int64(uid))
	response.Result(errCode, unReadNotice, context)
}

func (n *NoticeApi) Delete(context *gin.Context) {
	var param models.NoticeDeleteParam
	if err := context.ShouldBind(&param); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	errCode := n.noticeService.Delete(&param)
	response.Result(errCode, nil, context)
}

func (n *NoticeApi) GetList(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	if uid <= 0 {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	noticeList, errCode := n.noticeService.GetList(int64(uid))
	response.Result(errCode, noticeList, context)
}
