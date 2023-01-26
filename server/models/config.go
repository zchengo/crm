package models

type MailConfig struct {
	Id       int64  `gorm:"primaryKey"`
	Stmp     string `gorm:"stmp"`
	Port     int    `gorm:"port"`
	AuthCode string `gorm:"auth_code"`
	Email    string `gorm:"email"`
	Status   int    `gorm:"status"`
	Creator  int64  `gorm:"creator"`
	Created  int64  `gorm:"created"`
	Updated  int64  `gorm:"updated"`
}

type MailConfigSaveParam struct {
	Id       int64  `json:"id" binding:"omitempty,gt=0"`
	Stmp     string `json:"stmp" binding:"required,ip|hostname"`
	Port     int    `json:"port" binding:"required,gt=0"`
	AuthCode string `json:"authCode" binding:"required,gt=0"`
	Email    string `json:"email" binding:"required,email"`
	Status   int    `json:"status" binding:"omitempty,oneof=1 2"`
	Creator  int64  `json:"creator" binding:"omitempty"`
}

type MailConfigStatusParam struct {
	Id      int64 `json:"id" binding:"omitempty,gt=0"`
	Status  int   `json:"status" binding:"required,oneof=1 2"`
	Creator int64 `json:"creator" binding:"omitempty"`
}

type MailConfigDeleteParam struct {
	Id int64 `json:"id" binding:"required,gt=0"`
}

type MailConfigInfo struct {
	Id        int64  `json:"id"`
	Stmp      string `json:"stmp"`
	Port      int    `json:"port"`
	AuthCode  string `json:"authCode"`
	Email     string `json:"email"`
	Status    int    `json:"status"`
	Usability int    `json:"usability"`
}
