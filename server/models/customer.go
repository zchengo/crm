package models

type Customer struct {
	Id       int64  `gorm:"primaryKey"`
	Name     string `gorm:"name"`
	Source   string `gorm:"source"`
	Phone    string `gorm:"phone"`
	Email    string `gorm:"email"`
	Industry string `gorm:"industry"`
	Level    string `gorm:"level"`
	Remarks  string `gorm:"remarks"`
	Region   string `gorm:"region"`
	Address  string `gorm:"address"`
	Status   int    `gorm:"status"`
	Creator  int64  `gorm:"creator"`
	Created  int64  `gorm:"created"`
	Updated  int64  `gorm:"updated"`
}

type CustomerCreateParam struct {
	Name     string `json:"name,omitempty"`
	Source   string `json:"source,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Email    string `json:"email,omitempty"`
	Industry string `json:"industry,omitempty"`
	Level    string `json:"level,omitempty"`
	Remarks  string `json:"remarks,omitempty"`
	Region   string `json:"region,omitempty"`
	Address  string `json:"address,omitempty"`
	Status   int    `json:"status,omitempty"`
	Creator  int64  `json:"creator,omitempty"`
}

type CustomerUpdateParam struct {
	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Source   string `json:"source,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Email    string `json:"email,omitempty"`
	Industry string `json:"industry,omitempty"`
	Level    string `json:"level,omitempty"`
	Remarks  string `json:"remarks,omitempty"`
	Region   string `json:"region,omitempty"`
	Address  string `json:"address,omitempty"`
	Status   int    `json:"status,omitempty"`
}

type CustomerDeleteParam struct {
	Ids []int64 `json:"ids,omitempty"`
}

type CustomerQueryParam struct {
	Id      int64  `form:"id,omitempty"`
	Name    string `form:"name,omitempty"`
	Phone   string `form:"phone,omitempty"`
	Creator int64  `form:"creator"`
	Page    Page
}

type CustomerList struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Source   string `json:"source"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Industry string `json:"industry"`
	Level    string `json:"level"`
	Remarks  string `json:"remarks"`
	Region   string `json:"region"`
	Address  string `json:"address"`
	Status   int    `json:"status"`
	Created  int64  `json:"created"`
	Updated  int64  `json:"updated"`
}

type CustomerInfo struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Source   string `json:"source"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Industry string `json:"industry"`
	Level    string `json:"level"`
	Remarks  string `json:"remarks"`
	Region   string `json:"region"`
	Address  string `json:"address"`
	Status   int    `json:"status"`
}

type CustomerOption struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
