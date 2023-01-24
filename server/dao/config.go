package dao

import (
	"crm/global"
	"crm/models"
	"time"
)

type MailConfigDao struct {
}

func NewMailConfigDao() *MailConfigDao {
	return &MailConfigDao{}
}

func (m *MailConfigDao) Save(param *models.MailConfigSaveParam) error {
	if !isExists(param.Creator) {
		return create(param)
	}
	return update(param)
}

func (m *MailConfigDao) GetInfo(uid int64) (*models.MailConfig, error) {
	var mc models.MailConfig
	err := global.Db.Table(MAIL_CONFIG).Where("creator = ?", uid).First(&mc).Error
	if err != nil {
		return nil, err
	}
	return &mc, nil
}

func (m *MailConfigDao) UpdateStatus(param *models.MailConfigStatusParam) error {
	mailConfig := models.MailConfig{
		Status:  param.Status,
		Updated: time.Now().Unix(),
	}
	db := global.Db.Model(&mailConfig).Where("id = ? and creator = ?", param.Id, param.Creator)
	if err := db.Updates(&mailConfig).Error; err != nil {
		return err
	}
	return nil
}

func isExists(uid int64) bool {
	var mc models.MailConfig
	db := global.Db.Table(MAIL_CONFIG).Where("creator = ?", uid).First(&mc)
	return db.RowsAffected != NumberNull
}

func create(param *models.MailConfigSaveParam) error {
	mailConfig := models.MailConfig{
		Stmp:     param.Stmp,
		Port:     param.Port,
		AuthCode: param.AuthCode,
		Email:    param.Email,
		Status:   1,
		Creator:  param.Creator,
		Created:  time.Now().Unix(),
	}
	if err := global.Db.Create(&mailConfig).Error; err != nil {
		return err
	}
	return nil
}

func update(param *models.MailConfigSaveParam) error {
	mailConfig := models.MailConfig{
		Id:       param.Id,
		Stmp:     param.Stmp,
		Port:     param.Port,
		AuthCode: param.AuthCode,
		Email:    param.Email,
		Status:   param.Status,
		Creator:  param.Creator,
		Updated:  time.Now().Unix(),
	}
	db := global.Db.Model(&mailConfig).Select("*").Omit("id", "creator", "created")
	if err := db.Updates(&mailConfig).Error; err != nil {
		return err
	}
	return nil
}
