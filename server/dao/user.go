package dao

import (
	"crm/global"
	"crm/models"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (u *UserDao) Create(param *models.UserCreateParam) error {
	user := models.User{
		Email:    param.Email,
		Password: param.Password,
		Status:   1,
		Created:  time.Now().Unix(),
	}
	return global.Db.Create(&user).Error
}

func (u *UserDao) UpdatePass(email, password string) error {
	userPass := models.User{
		Password: password,
		Updated:  time.Now().Unix(),
	}
	return global.Db.Model(&models.User{}).Where("email = ?", email).Updates(&userPass).Error
}

func (u *UserDao) IsExists(email string) bool {
	rows := global.Db.Where("email = ?", email).First(&models.User{}).RowsAffected
	return rows != NumberNull
}

func (u *UserDao) GetUid(email string) (int64, error) {
	user, err := u.GetUser(email)
	if err != nil {
		return NumberNull, err
	}
	return user.Id, nil
}

func (u *UserDao) GetUser(email string) (*models.User, error) {
	var user models.User
	err := global.Db.Table(USER).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDao) DeleteData(param models.UserDeleteParam) error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		mw := map[interface{}]string{
			&models.Product{}:    "creator = ?",
			&models.Customer{}:   "creator = ?",
			&models.Contract{}:   "creator = ?",
			&models.MailConfig{}: "creator = ?",
			&models.Subscribe{}:  "uid = ?",
			&models.User{}:       "id = ?",
		}
		for k, v := range mw {
			if err := tx.Where(v, param.Id).Delete(k).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (u *UserDao) GetInfo(uid int64) (*models.UserPersonInfo, error) {
	var user models.UserPersonInfo
	err := global.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(USER).Where("id = ?", uid).First(&user).Error; err != nil {
			return err
		}
		var subscribe models.Subscribe
		if err := tx.Table(SUBSCRIBE).Select("version").Where("uid = ?", uid).First(&subscribe).Error; err != nil {
			return err
		}
		user.Version = subscribe.Version
		return nil
	})
	return &user, err
}

func (u *UserDao) SetCode(code int, email string) error {
	key := fmt.Sprintf("user:%s:code", email)
	return global.Rdb.SetEx(ctx, key, strconv.Itoa(code), 10*time.Minute).Err()
}

func (u *UserDao) GetCode(email string) string {
	key := fmt.Sprintf("user:%s:code", email)
	return global.Rdb.Get(ctx, key).Val()
}
