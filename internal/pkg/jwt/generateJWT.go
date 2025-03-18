package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("your-secret-key")

func GenerateJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,                            // 自定义字段，存储用户信息
		"exp":   time.Now().Add(time.Hour).Unix(), // 过期时间 1 小时
	})

	return token.SignedString(jwtSecret)
}
