package database

import "bookkeeping-server/internal/model"

func Migrate() {
	DB.AutoMigrate(&model.ValidationEmailCode{})
}
