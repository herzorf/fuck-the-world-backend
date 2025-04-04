package FTWJwt

import (
	"fmt"
	"fuck-the-world/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

var Secret = []byte("fuck-the-world")

func GenerateJWT(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 4).Unix(), // 过期时间 1 小时
	})

	return token.SignedString(Secret)
}

type JwtInfo struct {
	UserID   uint
	Role     string
	Username string
}

func ParseJWT(tokenString string) (JwtInfo, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil // 返回密钥
	}, jwt.WithoutClaimsValidation())
	var jInfo = JwtInfo{
		UserID:   0,
		Role:     "",
		Username: "",
	}
	if err != nil {
		return jInfo, fmt.Errorf("无效token")
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		userIDFloat, ok := (*claims)["userId"].(float64)
		username, ok := (*claims)["username"]
		if !ok {
			return jInfo, fmt.Errorf("userId 字段解析失败")
		}
		expFloat, ok := (*claims)["exp"].(float64)
		if !ok {
			return jInfo, fmt.Errorf("exp 字段解析失败")
		}
		role, ok := (*claims)["role"]
		if !ok {
			return jInfo, fmt.Errorf("role 字段解析失败")
		}
		expTime := time.Unix(int64(expFloat), 0)

		// 检查是否过期
		if time.Now().After(expTime) {
			return jInfo, fmt.Errorf("token 已过期,请重新登录")
		}
		jInfo.UserID = uint(userIDFloat)
		jInfo.Role = role.(string)
		jInfo.Username = username.(string)
		return jInfo, nil
	} else {
		log.Println("token无效载体", token.Valid)
	}

	return jInfo, fmt.Errorf("token 无效")
}
