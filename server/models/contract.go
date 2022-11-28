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
	Name        string       `json:"name,omitempty"`
	Amount      float64      `json:"amount,omitempty"`
	BeginTime   string       `json:"beginTime,omitempty"`
	OverTime    string       `json:"overTime,omitempty"`
	Remarks     string       `json:"remarks,omitempty"`
	Cid         int64        `json:"cid,omitempty"`
	Productlist *Productlist `json:"productlist,omitempty"`
	Status      int          `json:"status,omitempty"`
	Creator     int64        `json:"creator,omitempty"`
}

type ContractUpdateParam struct {
	Id          int64        `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Amount      float64      `json:"amount,omitempty"`
	BeginTime   string       `json:"beginTime,omitempty"`
	OverTime    string       `json:"overTime,omitempty"`
	Remarks     string       `json:"remarks,omitempty"`
	Cid         int64        `json:"cid,omitempty"`
	Productlist *Productlist `json:"productlist,omitempty"`
	Status      int          `json:"status,omitempty"`
}

type ContractDeleteParam struct {
	Ids []int64 `json:"ids,omitempty"`
}

type ContractQueryParam struct {
	Id      int64  `form:"id,omitempty"`
	Name    string `form:"name,omitempty"`
	Creator int64  `form:"creator,omitempty"`
	Page    Page
}

type ContractList struct {
	Id        int64   `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	Amount    float64 `json:"amount,omitempty"`
	BeginTime string  `json:"beginTime,omitempty"`
	OverTime  string  `json:"overTime,omitempty"`
	Remarks   string  `json:"remarks,omitempty"`
	Cname     string  `json:"cname,omitempty"`
	Status    int     `json:"status,omitempty"`
	Created   int64   `json:"created,omitempty"`
	Updated   int64   `json:"updated,omitempty"`
}

type ContractInfo struct {
	Id          int64        `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Cid         int64        `json:"cid,omitempty"`
	Amount      float64      `json:"amount,omitempty"`
	BeginTime   string       `json:"beginTime,omitempty"`
	OverTime    string       `json:"overTime,omitempty"`
	Remarks     string       `json:"remarks,omitempty"`
	Productlist *Productlist `json:"productlist,omitempty"`
	Status      int          `json:"status,omitempty"`
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

type Productlist []*Products

func (p *Productlist) Value() (driver.Value, error) {
	b, err := json.Marshal(p)
	return string(b), err
}

func (p *Productlist) Scan(src any) error {
	return json.Unmarshal(src.([]byte), p)
}
