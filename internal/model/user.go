package model

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Email string `gorm:"size:255;not null"`
}
