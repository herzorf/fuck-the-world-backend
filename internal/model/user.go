package model

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"size:255;not null"`
	Password  string `gorm:"size:255;not null"`
	Email     string `gorm:"size:255;not null"`
	Role      string `gorm:"size:50;not null;default:operator"` // 角色字段
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 角色枚举
const (
	RoleAdmin    = "admin"
	RoleOperator = "operator"
)
