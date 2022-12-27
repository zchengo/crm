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
	Name        string  `json:"name" binding:"required"`
	Type        int     `json:"type" binding:"required,len=1"`
	Unit        string  `json:"unit" binding:"-"`
	Code        string  `json:"code" binding:"-"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Description string  `json:"description" binding:"-"`
	Status      int     `json:"status" binding:"required,oneof=1 2"`
	Creator     int64   `json:"creator,omitempty" binding:"-"`
}

type ProductUpdateParam struct {
	Id          int64   `json:"id" binding:"required,gt=0"`
	Name        string  `json:"name" binding:"required"`
	Type        int     `json:"type" binding:"required,len=1"`
	Unit        string  `json:"unit" binding:"-"`
	Code        string  `json:"code" binding:"-"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Description string  `json:"description" binding:"-"`
	Status      int     `json:"status" binding:"required,oneof=1 2"`
	Creator     int64   `json:"creator,omitempty" binding:"-"`
}

type ProductDeleteParam struct {
	Ids []int64 `json:"ids" binding:"required"`
}

type ProductQueryParam struct {
	Id      int64   `form:"id" binding:"omitempty,gt=0"`
	Ids     []int64 `form:"ids" json:"ids" binding:"-"`
	Name    string  `form:"name" binding:"-"`
	Status  int     `form:"status" binding:"omitempty,oneof=1 2"`
	Creator int64   `form:"creator,omitempty" binding:"-"`
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

type ProductExcelRow struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Unit        string  `json:"unit"`
	Code        string  `json:"code"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	Created     string  `json:"created"`
	Updated     string  `json:"updated"`
}
