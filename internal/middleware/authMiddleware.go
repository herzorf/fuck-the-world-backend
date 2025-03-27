package middleware

import (
	FTWJwj "fuck-the-world/internal/pkg/jwt"
	"fuck-the-world/unit"
	"github.com/gin-gonic/gin"
	"log"
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
		userId, err := FTWJwj.ParseJWT(tokenString)
		if err != nil {
			unit.RespondJSON(c, http.StatusUnauthorized, err.Error(), nil)
			c.Abort()
			return
		}
		// 检查 Token 是否过期
		exp, ok := claims["exp"].(float64)
		if !ok || int64(exp) < time.Now().Unix() {
			unit.RespondJSON(c, http.StatusUnauthorized, "Token 已过期", nil)
			c.Abort()
			return
		}

		// 获取 userId
		userId, ok := claims["userId"].(float64)
		if !ok {
			unit.RespondJSON(c, http.StatusUnauthorized, "Token 缺少用户信息", nil)
			c.Abort()
			return
		}

		// 将 userId 存入请求上下文，供后续处理使用
		c.Set("userId", uint(userId))
		// 将 role 存入请求上下文，供后续处理使用
		c.Set("role", claims["role"])
		// 继续执行下一个中间件或请求处理
		c.Next()
	}
}
