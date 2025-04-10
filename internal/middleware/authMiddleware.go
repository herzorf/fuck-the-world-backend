package middleware

import (
	FTWJwj "fuck-the-world/internal/pkg/jwt"
	"fuck-the-world/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization 头部
		tokenString := c.GetHeader("Authorization")
		if len(tokenString) == 0 {
			utils.RespondJSON(c, http.StatusUnauthorized, "未登陆", nil)
			c.Abort()
			return
		}
		jInfo, err := FTWJwj.ParseJWT(tokenString)
		if err != nil {
			utils.RespondJSON(c, http.StatusUnauthorized, err.Error(), nil)
			c.Abort()
			return
		}
		// 将 userId 存入请求上下文，供后续处理使用
		c.Set("userId", jInfo.UserID)
		// 将username 存到请求上下文中，供后续使用
		c.Set("username", jInfo.Username)
		// 将 role 存入请求上下文，供后续处理使用
		c.Set("role", jInfo.Role)
		// 继续执行下一个中间件或请求处理
		c.Next()
	}
}
