# 用户授权与认证

Crm 系统采用 JWT 来完成用户授权与认证。

## 什么是 JWT ？

JSON Web Token（JWT）是一个开放标准（RFC 7519），它定义了一种紧凑和自包含的方式，用于作为JSON对象在各方之间安全地传输信息。

这些信息可以被验证和信任，因为它是数字签名的。JWT可以使用秘密（使用HMAC算法）或使用RSA或ECDSA的公钥/私钥对进行签名。

以下是一些 JSON Web Token 应用场景：

- 授权：这是使用 JWT 最常见的场景。一旦用户登录，每个后续请求都将包括 JWT，允许用户访问该令牌允许的路由、服务和资源。单点登录是当今广泛使用 JWT 的一项功能，因为它的开销很小，并且可以轻松地在不同领域使用。

- 信息交换：JWT 是各方之间安全传输信息的好方法。因为JWT可以签名，例如，使用公钥或私钥对，您可以确定发件人就是他们所说的那个人。此外，由于签名是使用标头和有效负载计算的，您还可以验证内容是否未被篡改。

想了解有关 JWT 的更多信息，请访问[jwt.io](https://jwt.io)。

## 安装

添加 ```golang-jwt``` 作为Go程序中的依赖项，请执行诶下命令：

```bash
go get -u github.com/golang-jwt/jwt/v4
```

## 封装成方法

请参考如下代码：

```go
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
```

想要了解有关 golang-jwt 的更多信息，请访问[golang-jwt](https://github.com/golang-jwt/jwt)。

## 当用户登录成功后生成 Token 信息

请参考 ```crm/server/service/user.go``` 中的 ```Login``` 方法，部分代码如下：

```go
// 用户登录
func (u *UserService) Login(param *models.UserLoginParam) (*models.UserInfo, int) {
    
    ...

    // 生成Token
	token, err := common.GenToken(user.Id)
	if err != nil {
		log.Printf("[error]Login:GenerateToken:%s", err)
		return nil, response.ErrCodeFailed
	}

    userInfo := models.UserInfo{
		Uid:   user.Id,
		Token: token,
	}

    ...
}
```

## 使用 JWT 中间件校验 Token

请参考如下代码：

```go
// JWT认证中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := strconv.Atoi(c.Request.Header.Get("uid"))
		token := c.Request.Header.Get("token")
		if token == "" {
			response.Result(response.ErrCodeNoLogin, nil, c)
			c.Abort()
			return
		}
		userid, err := common.VerifyToken(token)
		if userid != int64(uid) || err != nil {
			response.Result(response.ErrCodeTokenExpire, nil, c)
			c.Abort()
			return
		}
		c.Next()
	}
}
```

前端发送请求时会携带 ```uid``` 和 ```token```，JWT 认证中间件会校验 ```token``` 的合法性来判断用户是否登录。

