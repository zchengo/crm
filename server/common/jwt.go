package common

import (
	"crm/global"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Uid int64 `json:"uid"`
	jwt.RegisteredClaims
}

// 生成Token
func GenToken(uid int64) (string, error) {
	var signingKey = []byte(global.Config.Jwt.SigningKey)
	var expiredTime = global.Config.Jwt.ExpiredTime
	claims := Claims{uid, jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Duration(expiredTime) * time.Second)},
		Issuer:    "crm",
	}}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(signingKey)
	return token, err
}

// 校验Token
func VerifyToken(tokens string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokens, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.SigningKey), nil
	})
	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		return claims.Uid, nil
	}
	return 0, err
}
