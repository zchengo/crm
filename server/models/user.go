package models

type User struct {
	Id       int64  `gorm:"primaryKey"`
	Email    string `gorm:"email"`
	Password string `gorm:"password"`
	Name     string `gorm:"name"`
	Version  int    `gorm:"version"`
	Expired  int64  `gorm:"expired"`
	Status   int    `gorm:"status"`
	Created  int64  `gorm:"created"`
	Updated  int64  `gorm:"updated"`
}

type UserCreateParam struct {
	Email    string `json:"email"`
	Code     string `json:"code"`
	Password string `json:"password"`
}

type UserUpdateParam struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   int    `json:"status"`
}

type UserDeleteParam struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
	Code  string `json:"code"`
}

type UserLoginParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserPassParam struct {
	Email    string `json:"email"`
	Code     string `json:"code"`
	Password string `json:"password"`
}

type UserMailParam struct {
	Email    string `json:"email"`
	Code     string `json:"code"`
	NewEmail string `json:"newEmail"`
}

type UserInfo struct {
	Uid   int64  `json:"uid"`
	Ver   int    `json:"version"`
	Token string `json:"token"`
}

type UserPersonInfo struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Version int    `json:"version"`
	Expired int64  `json:"expired"`
}

type UserVerisonInfo struct {
	Version int   `json:"version"`
	Expired int64 `json:"expired"`
}
