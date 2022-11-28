package common

import (
	"crm/global"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var signingKey = []byte(global.Config.Jwt.SigningKey)

// 生成Token
func GenToken(uid int64, expire int64) (string, error) {
	username := strconv.FormatInt(uid, 20)
	type Claims struct {
		Username string `json:"username"`
		jwt.RegisteredClaims
	}
	claims := Claims{username, jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Duration(expire) * time.Hour)},
		Issuer:    username,
	}}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(signingKey)
	return token, err
}
