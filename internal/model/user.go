package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey;comment:'用户ID'"`
	Username  string    `gorm:"size:255;not null;comment:'用户名'"`
	Password  string    `gorm:"size:255;not null;comment:'密码'"`
	Role      string    `gorm:"size:50;not null;default:'operator';comment:'角色'"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:'创建时间'"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:'更新时间'"`
	IsActive  bool      `gorm:"default:true;comment:'是否可用'"`
	IsDeleted bool      `gorm:"default:false;comment:'是否删除'"`
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
