package service

import (
	"crm/common"
	"crm/dao"
	"crm/models"
	"crm/response"
)

type MailConfigService struct {
	mailConfigDao *dao.MailConfigDao
}

func NewMailConfigService() *MailConfigService {
	return &MailConfigService{
		mailConfigDao: dao.NewMailConfigDao(),
	}
}

// 保存邮件服务配置
func (m *MailConfigService) Save(param *models.MailConfigSaveParam) int {
	if err := m.mailConfigDao.Save(param); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 删除邮件服务配置
func (m *MailConfigService) Delete(param *models.MailConfigDeleteParam) int {
	if err := m.mailConfigDao.Delete(param); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 开启或关闭邮件服务
func (m *MailConfigService) UpdateStatus(param *models.MailConfigStatusParam) int {
	if !m.mailConfigDao.IsExists(param.Creator) {
		return response.ErrCodeMailConfigUnSet
	}
	if err := m.mailConfigDao.UpdateStatus(param); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 获取邮件配置信息
func (m *MailConfigService) GetInfo(uid int64) (*models.MailConfigInfo, int) {
	if !m.mailConfigDao.IsExists(uid) {
		return nil, response.ErrCodeMailConfigUnSet
	}
	mc, err := m.mailConfigDao.GetInfo(uid)
	if err != nil {
		return nil, response.ErrCodeFailed
	}
	mailConfig := models.MailConfigInfo{
		Id:       mc.Id,
		Stmp:     mc.Stmp,
		Port:     mc.Port,
		AuthCode: mc.AuthCode,
		Email:    mc.Email,
		Status:   mc.Status,
	}
	return &mailConfig, response.ErrCodeSuccess
}

// 检查邮件配置的有效性
func (m *MailConfigService) Check(uid int64) int {
	mc, err := m.mailConfigDao.GetInfo(uid)
	if err != nil {
		return response.ErrCodeMailConfigInvalid
	}
	if err = common.DialMail(mc.Stmp, mc.Port, mc.Email, mc.AuthCode); err != nil {
		return response.ErrCodeMailConfigInvalid
	}
	return response.ErrCodeSuccess
}
