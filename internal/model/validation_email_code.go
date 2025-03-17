package model

import "time"

type ValidationEmailCode struct {
	ID        uint       `gorm:"primaryKey"`
	Code      string     `gorm:"size:20;not null"`
	Email     string     `gorm:"size:255;not null"`
	UsedAt    *time.Time // 可以为空，代表未使用
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}
