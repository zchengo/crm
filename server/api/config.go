package api

import (
	"crm/models"
	"crm/response"
	"crm/service"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MailConfigApi struct {
	mailConfigService *service.MailConfigService
}

func NewMailConfigApi() *MailConfigApi {
	return &MailConfigApi{
		mailConfigService: service.NewMailConfigService(),
	}
}

// 保存邮件服务配置
func (m *MailConfigApi) Save(context *gin.Context) {
	var param models.MailConfigSaveParam
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	err := context.ShouldBind(&param)
	if uid <= 0 || err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		log.Println(err)
		return
	}
	param.Creator = int64(uid)
	errCode := m.mailConfigService.Save(&param)
	if errCode == 0 {
		mailConfig, errCode := m.mailConfigService.GetInfo(int64(uid))
		response.Result(errCode, mailConfig, context)
		return
	}
	response.Result(errCode, nil, context)
}

// 删除邮件服务配置
func (m *MailConfigApi) Delete(context *gin.Context) {
	var param models.MailConfigDeleteParam
	if err := context.ShouldBind(&param); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		log.Println(err)
		return
	}
	errCode := m.mailConfigService.Delete(&param)
	response.Result(errCode, nil, context)
}

// 开启或关闭邮件服务
func (m *MailConfigApi) UpdateStatus(context *gin.Context) {
	var param models.MailConfigStatusParam
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	err := context.ShouldBind(&param)
	if uid <= 0 || err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		log.Println(err)
		return
	}
	param.Creator = int64(uid)
	errCode := m.mailConfigService.UpdateStatus(&param)
	response.Result(errCode, nil, context)
}

// 获取邮件服务配置信息
func (m *MailConfigApi) GetInfo(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	if uid <= 0 {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	mailConfig, errCode := m.mailConfigService.GetInfo(int64(uid))
	response.Result(errCode, mailConfig, context)
}

// 检查邮件配置的有效性
func (m *MailConfigApi) Check(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	if uid <= 0 {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	errCode := m.mailConfigService.Check(int64(uid))
	response.Result(errCode, nil, context)
}