package models

type Subscribe struct {
	Id      int64 `gorm:"primaryKey"`
	Uid     int64 `gorm:"uid"`
	Version int   `gorm:"version"`
	Expired int64 `gorm:"expired"`
	Created int64 `gorm:"created"`
	Updated int64 `gorm:"updated"`
}

type SubscribeCreateParam struct {
	Uid     int64 `json:"uid"`
	Version int   `json:"version"`
	Expired int64 `json:"expired"`
}

type SubscribeUpdateParam struct {
	Id      int64 `json:"id"`
	Uid     int64 `json:"uid"`
	Version int   `json:"version"`
	Expired int64 `json:"expired"`
}

type SubscribePayParam struct {
	Uid      int64 `json:"uid" binding:"-"`
	Duration int64 `json:"duration" binding:"required,oneof=30 90 180 365 730"`
}

type SubscribePayOrder struct {
	Uid      int64  `json:"uid"`
	TradeNo  string `json:"tradeNo"`
	Duration int64  `json:"duration"`
}

type SubscribePayUrl struct {
	PayUrl string `json:"payUrl"`
}

type SubscribeInfo struct {
	Version int   `json:"version"`
	Expired int64 `json:"expired"`
}
