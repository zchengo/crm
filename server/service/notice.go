package service

import (
	"crm/dao"
	"crm/models"
	"crm/response"
)

type NoticeService struct {
	noticeDao *dao.NoticeDao
}

func NewNoticeService() *NoticeService {
	return &NoticeService{
		noticeDao: dao.NewNoticeDao(),
	}
}

// 新建通知
func (n *NoticeService) Create(param *models.NoticeCreateParam) int {
	if err := n.noticeDao.Create(param); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 更新通知
func (n *NoticeService) Update(param *models.NoticeUpdateParam) int {
	if err := n.noticeDao.Update(param); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 获取未读通知数量
func (n *NoticeService) GetUnReadCount(uid int64) (models.UnReadNotice, int) {
	unReadCount, err := n.noticeDao.GetUnReadCountByUid(uid)
	if err != nil {
		return unReadCount, response.ErrCodeFailed
	}
	return unReadCount, response.ErrCodeSuccess
}

// 删除通知
func (n *NoticeService) Delete(param *models.NoticeDeleteParam) int {
	if err := n.noticeDao.Delete(param); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

// 获取通知列表
func (n *NoticeService) GetList(uid int64) ([]*models.NoticeList, int) {
	noticeList, err := n.noticeDao.GetList(uid)
	if err != nil {
		return nil, response.ErrCodeFailed
	}
	return noticeList, response.ErrCodeSuccess
}
