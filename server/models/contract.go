package models

import (
	"database/sql/driver"
	"encoding/json"
)

type Contract struct {
	Id          int64        `gorm:"primaryKey"`
	Name        string       `gorm:"name"`
	Amount      float64      `gorm:"amount"`
	BeginTime   string       `gorm:"begin_time"`
	OverTime    string       `gorm:"over_time"`
	Remarks     string       `gorm:"remarks"`
	Cid         int64        `gorm:"cid"`
	Productlist *Productlist `gorm:"type:json"`
	Status      int          `gorm:"status"`
	Creator     int64        `gorm:"creator"`
	Created     int64        `gorm:"created"`
	Updated     int64        `gorm:"updated"`
}

type ContractCreateParam struct {
	Name        string       `json:"name" binding:"required"`
	Amount      float64      `json:"amount" binding:"required,gt=0"`
	BeginTime   string       `json:"beginTime" binding:"-"`
	OverTime    string       `json:"overTime" binding:"-"`
	Remarks     string       `json:"remarks" binding:"-"`
	Cid         int64        `json:"cid" binding:"required,gt=0"`
	Productlist *Productlist `json:"productlist"`
	Status      int          `json:"status" binding:"required,oneof=1 2"`
	Creator     int64        `json:"creator,omitempty" binding:"-"`
}

type ContractUpdateParam struct {
	Id          int64        `json:"id" binding:"required,gt=0"`
	Name        string       `json:"name" binding:"required"`
	Amount      float64      `json:"amount" binding:"omitempty,gt=0"`
	BeginTime   string       `json:"beginTime" binding:"-"`
	OverTime    string       `json:"overTime" binding:"-"`
	Remarks     string       `json:"remarks" binding:"-"`
	Cid         int64        `json:"cid" binding:"required,gt=0"`
	Productlist *Productlist `json:"productlist"`
	Status      int          `json:"status" binding:"required,oneof=1 2"`
}

type ContractDeleteParam struct {
	Ids []int64 `json:"ids" binding:"required"`
}

type ContractQueryParam struct {
	Id      int64   `form:"id" binding:"omitempty,gt=0"`
	Pids    []int64 `form:"pids" json:"pids" binding:"-"`
	Name    string  `form:"name" binding:"-"`
	Status  int     `form:"status" binding:"omitempty,oneof=1 2"`
	Creator int64   `form:"creator,omitempty" binding:"-"`
	Page    Page
}

type ContractList struct {
	Id        int64   `json:"id"`
	Name      string  `json:"name"`
	Amount    float64 `json:"amount"`
	BeginTime string  `json:"beginTime"`
	OverTime  string  `json:"overTime"`
	Remarks   string  `json:"remarks"`
	Cname     string  `json:"cname"`
	Status    int     `json:"status"`
	Created   int64   `json:"created"`
	Updated   int64   `json:"updated"`
}

type ContractInfo struct {
	Id          int64        `json:"id"`
	Name        string       `json:"name"`
	Cid         int64        `json:"cid"`
	Amount      float64      `json:"amount"`
	BeginTime   string       `json:"beginTime"`
	OverTime    string       `json:"overTime"`
	Remarks     string       `json:"remarks"`
	Productlist *Productlist `json:"productlist"`
	Status      int          `json:"status"`
}

type Products struct {
	Id    int64   `json:"id"`
	Name  string  `json:"name"`
	Type  int     `json:"type"`
	Unit  string  `json:"unit"`
	Price float64 `json:"price"`
	Count int     `json:"count"`
	Total float64 `json:"total"`
}

type ContractExcelRow struct {
	Name      string  `json:"name"`
	Cname     string  `json:"cname"`
	Amount    float64 `json:"amount"`
	BeginTime string  `json:"beginTime"`
	OverTime  string  `json:"overTime"`
	Remarks   string  `json:"remarks"`
	Status    string  `json:"status"`
	Created   string  `json:"created"`
	Updated   string  `json:"updated"`
}

type Productlist []*Products

func (p *Productlist) Value() (driver.Value, error) {
	b, err := json.Marshal(p)
	return string(b), err
}

func (p *Productlist) Scan(src any) error {
	return json.Unmarshal(src.([]byte), p)
}
