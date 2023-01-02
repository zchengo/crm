package service

import (
	"crm/global"
	"crm/models"
	"crm/response"
	"time"
)

const (
	Read   = 1 // 已读
	UnRead = 2 // 未读
)

type NoticeService struct {
}

// 新建通知
func (n *NoticeService) Create(param *models.NoticeParam) int {
	notice := models.Notice{
		Content: param.Content,
		Status:  UnRead,
		Creator: param.Creator,
		Created: time.Now().Unix(),
	}
	if err := global.Db.Create(&notice).Error; err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 更新通知
func (n *NoticeService) Update(param *models.NoticeUpdateParam) int {
	notice := models.Notice{
		Id:      param.Id,
		Status:  Read,
		Updated: time.Now().Unix(),
	}
	if err := global.Db.Model(&notice).Updates(&notice).Error; err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 获取未读通知数量
func (n *NoticeService) UnReadCount(uid int64) (models.UnReadNotice, int) {
	var unRead models.UnReadNotice
	raw := "select count(*) from notice where status = 2 and creator = ?"
	if err := global.Db.Raw(raw, uid).Scan(&unRead.Count).Error; err != nil {
		return unRead, response.ErrCodeFailed
	}
	return unRead, response.ErrCodeSuccess
}

// 删除通知
func (n *NoticeService) Delete(param *models.NoticeDeleteParam) int {
	err := global.Db.Delete(&models.Notice{}, param.Ids).Error
	if err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 获取通知列表
func (n *NoticeService) GetList(uid int64) ([]*models.NoticeList, int) {
	noticeList := make([]*models.NoticeList, 0)
	if err := global.Db.Table(NOTICE).Where("creator = ?", uid).Order("status desc").Order("created desc").Find(&noticeList).Error; err != nil {
		return nil, response.ErrCodeFailed
	}
	return noticeList, response.ErrCodeSuccess
}
