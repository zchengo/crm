package dao

import (
	"crm/global"
	"crm/models"
	"time"
)

const (
	Read   = 1 // 已读
	UnRead = 2 // 未读
)

type NoticeDao struct {
}

func NewNoticeDao() *NoticeDao {
	return &NoticeDao{}
}

func (n *NoticeDao) Create(param *models.NoticeCreateParam) error {
	notice := models.Notice{
		Content: param.Content,
		Status:  UnRead,
		Creator: param.Creator,
		Created: time.Now().Unix(),
	}
	return global.Db.Create(&notice).Error
}

func (n *NoticeDao) Update(param *models.NoticeUpdateParam) error {
	notice := models.Notice{
		Id:      param.Id,
		Status:  Read,
		Updated: time.Now().Unix(),
	}
	return global.Db.Model(&notice).Updates(&notice).Error
}

func (n *NoticeDao) GetUnReadCountByUid(uid int64) (models.UnReadNotice, error) {
	var unRead models.UnReadNotice
	raw := "select count(*) from notice where status = 2 and creator = ?"
	if err := global.Db.Raw(raw, uid).Scan(&unRead.Count).Error; err != nil {
		return unRead, err
	}
	return unRead, nil
}

func (n *NoticeDao) Delete(param *models.NoticeDeleteParam) error {
	return global.Db.Delete(&models.Notice{}, param.Ids).Error
}

func (n *NoticeDao) GetList(uid int64) ([]*models.NoticeList, error) {
	noticeList := make([]*models.NoticeList, 0)
	db := global.Db.Table(NOTICE).Where("creator = ?", uid)
	if err := db.Order("status desc").Order("created desc").Find(&noticeList).Error; err != nil {
		return nil, err
	}
	return noticeList, nil
}
