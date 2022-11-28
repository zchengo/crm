package models

type Product struct {
	Id          int64   `gorm:"primaryKey"`
	Name        string  `gorm:"name"`
	Type        int     `gorm:"type"`
	Unit        string  `gorm:"unit"`
	Code        string  `gorm:"code"`
	Price       float64 `gorm:"price"`
	Description string  `gorm:"description"`
	Status      int     `gorm:"status"`
	Creator     int64   `gorm:"creator"`
	Created     int64   `gorm:"created"`
	Updated     int64   `gorm:"updated"`
}

type ProductCreateParam struct {
	Name        string  `json:"name,omitempty"`
	Type        int     `json:"type,omitempty"`
	Unit        string  `json:"unit,omitempty"`
	Code        string  `json:"code,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Description string  `json:"description,omitempty"`
	Status      int     `json:"status,omitempty"`
	Creator     int64   `json:"creator,omitempty"`
}

type ProductUpdateParam struct {
	Id          int64   `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Type        int     `json:"type,omitempty"`
	Unit        string  `json:"unit,omitempty"`
	Code        string  `json:"code,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Description string  `json:"description,omitempty"`
	Status      int     `json:"status,omitempty"`
	Creator     int64   `json:"creator,omitempty"`
}

type ProductDeleteParam struct {
	Ids []int64 `json:"ids,omitempty"`
}

type ProductQueryParam struct {
	Id      int64   `form:"id,omitempty"`
	Ids     []int64 `form:"ids,omitempty"`
	Name    string  `form:"name,omitempty"`
	Status  int     `form:"status,omitempty"`
	Creator int64   `form:"creator,omitempty"`
	Page    Page
}

type ProductList struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Type        int     `json:"type"`
	Unit        string  `json:"unit"`
	Code        string  `json:"code"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Status      int     `json:"status"`
	Created     int64   `json:"created"`
	Updated     int64   `json:"updated"`
}

type ProductInfo struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Type        int     `json:"type"`
	Unit        string  `json:"unit"`
	Code        string  `json:"code"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Status      int     `json:"status"`
}
