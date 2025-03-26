package FTWJwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var Secret = []byte("fuck-the-world")

func GenerateJWT(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": id,                               // 自定义字段，存储用户信息
		"exp":    time.Now().Add(time.Hour).Unix(), // 过期时间 1 小时
	})

	return token.SignedString(Secret)
}
