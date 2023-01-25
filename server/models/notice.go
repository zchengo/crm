package models

type Notice struct {
	Id      int64  `gorm:"primaryKey"`
	Content string `gorm:"content"`
	Status  int    `gorm:"status"`
	Creator int64  `gorm:"creator"`
	Created int64  `gorm:"created"`
	Updated int64  `gorm:"updated"`
}

type NoticeCreateParam struct {
	Content string `json:"content"`
	Creator int64  `gorm:"creator"`
}

type NoticeUpdateParam struct {
	Id int64 `json:"id" binding:"required,gt=0"`
}

type NoticeDeleteParam struct {
	Ids []int64 `json:"ids" binding:"required"`
}

type UnReadNotice struct {
	Count int `json:"count"`
}

type NoticeList struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
	Status  int    `json:"status"`
	Created int64  `json:"created"`
}
