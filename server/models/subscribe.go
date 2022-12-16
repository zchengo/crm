package models

type Subscribe struct {
	Id      int64 `gorm:"primaryKey"`
	Uid     int64 `gorm:"uid"`
	Version int   `gorm:"version"`
	Expired int64 `gorm:"expired"`
	Created int64 `gorm:"created"`
	Updated int64 `gorm:"updated"`
}

type SubscribePayParam struct {
	Uid     int64 `json:"uid" binding:"-"`
	Version int   `json:"version" binding:"required,oneof=2 3"`
}

type SubscribePayOrder struct {
	Uid     int64 `json:"uid"`
	Version int   `json:"version"`
}

type SubscribePayUrl struct {
	PayUrl string `json:"payUrl"`
}

type SubscribeInfo struct {
	Version int   `json:"version"`
	Expired int64 `json:"expired"`
}
