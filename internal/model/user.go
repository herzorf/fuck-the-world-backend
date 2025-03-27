package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"size:255;not null"`
	Password  string `gorm:"size:255;not null"`
	Role      string `gorm:"size:50;not null;default:operator"` // 角色字段
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 角色枚举
const (
	RoleAdmin    = "admin"
	RoleOperator = "operator"
)

// HashPassword 加密密码
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword 验证密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
