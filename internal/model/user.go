package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;comment:'用户ID'"`
	Username  string    `json:"username" gorm:"size:255;not null;comment:'用户名'"`
	Password  string    `json:"password" gorm:"size:255;not null;comment:'密码'"`
	Role      string    `json:"role" gorm:"size:50;not null;default:'operator';comment:'角色'"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'创建时间'"`                             // 确保格式
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'更新时间'"` // 自动更新
	IsActive  *bool     `json:"isActive" gorm:"default:true;comment:'是否可用'"`
	IsDeleted *bool     `json:"isDeleted" gorm:"default:false;comment:'是否删除'"`
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
