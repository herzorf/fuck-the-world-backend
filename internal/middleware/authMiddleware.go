package middleware

import (
	FTWJwj "fuck-the-world/internal/pkg/jwt"
	"fuck-the-world/unit"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization 头部
		tokenString := c.GetHeader("Authorization")
		if len(tokenString) == 0 {
			unit.RespondJSON(c, http.StatusUnauthorized, "未登陆", nil)
			c.Abort()
			return
		}
		jInfo, err := FTWJwj.ParseJWT(tokenString)
		if err != nil {
			unit.RespondJSON(c, http.StatusUnauthorized, err.Error(), nil)
			c.Abort()
			return
		}
		// 将 userId 存入请求上下文，供后续处理使用
		c.Set("userId", jInfo.UserID)
		// 将 role 存入请求上下文，供后续处理使用
		c.Set("role", jInfo.Role)
		// 继续执行下一个中间件或请求处理
		c.Next()
	}
}
