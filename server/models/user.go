package models

type User struct {
	Id       int64  `gorm:"primaryKey"`
	Email    string `gorm:"email"`
	Password string `gorm:"password"`
	Name     string `gorm:"name"`
	Status   int    `gorm:"status"`
	Created  int64  `gorm:"created"`
	Updated  int64  `gorm:"updated"`
}

type UserCreateParam struct {
	Email    string `json:"email" binding:"required,email"`
	Code     string `json:"code" binding:"required,len=6"`
	Password string `json:"password" binding:"required"`
}

type UserDeleteParam struct {
	Id    int64  `json:"id,omitempty" binding:"-"`
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required,len=6"`
}

type UserLoginParam struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserVerifyCodeParam struct {
	Email string `form:"email" binding:"required,email"`
}

type UserPassParam struct {
	Email    string `json:"email" binding:"required,email"`
	Code     string `json:"code" binding:"required,len=6"`
	Password string `json:"password" binding:"required"`
}

type UserMailParam struct {
	Email    string `json:"email" binding:"required,email"`
	Code     string `json:"code" binding:"required,len=6"`
	NewEmail string `json:"newEmail" binding:"required,email"`
}

type UserInfo struct {
	Uid   int64  `json:"uid"`
	Token string `json:"token"`
}

type UserPersonInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
