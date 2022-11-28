package middleware

import (
	"crm/response"
	"crm/service"

	"github.com/gin-gonic/gin"
)

// JwtAuth JWT认证中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.Result(response.ErrCodeNoLogin, nil, c)
			c.Abort()
			return
		}
		user := &service.UserService{}
		if err := user.VerifyToken(token); err != nil {
			response.Result(response.ErrCodeTokenExpire, nil, c)
			c.Abort()
			return
		}
		c.Next()
	}
}
