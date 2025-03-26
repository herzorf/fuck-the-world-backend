package FTWJwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
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
func ParseJWT(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil // 返回密钥
	}, jwt.WithoutClaimsValidation())

	if err != nil {
		return 0, fmt.Errorf("无效token")
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		userIDFloat, ok := (*claims)["userId"].(float64)
		if !ok {
			return 0, fmt.Errorf("userId 字段解析失败")
		}
		expFloat, ok := (*claims)["exp"].(float64)
		if !ok {
			return 0, fmt.Errorf("exp 字段解析失败")
		}
		expTime := time.Unix(int64(expFloat), 0)

		// 检查是否过期
		if time.Now().After(expTime) {
			return 0, fmt.Errorf("token 已过期")
		}

		return uint(userIDFloat), nil
	} else {
		log.Println("token无效载体", token.Valid)
	}

	return 0, fmt.Errorf("token 无效")
}
