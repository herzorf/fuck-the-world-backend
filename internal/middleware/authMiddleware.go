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
		log.Println("userId", userId)
		c.Set("userId", userId)
		c.Next()
	}
}
