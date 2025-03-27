package FTWJwt

import (
	"fuck-the-world/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var Secret = []byte("fuck-the-world")

func GenerateJWT(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"role":   user.Role,
		"exp":    time.Now().Add(time.Hour).Unix(), // 过期时间 1 小时
	})

	return token.SignedString(Secret)
}
