package middleware

import (
	"errors"
	FTWJwj "fuck-the-world/internal/pkg/jwt"
	"fuck-the-world/unit"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization 头部
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			unit.RespondJSON(c, http.StatusUnauthorized, "未登陆", nil)
			c.Abort()
			return
		}

		// 去掉 "Bearer " 前缀
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		// 解析 JWT
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("签名方法不匹配")
			}
			return FTWJwj.Secret, nil
		})
		log.Println("44444444", token.Valid)
		// 检查解析是否成功
		if err != nil || !token.Valid {
			unit.RespondJSON(c, http.StatusUnauthorized, "无效Token", nil)
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

		// 继续执行下一个中间件或请求处理
		c.Next()
	}
}
