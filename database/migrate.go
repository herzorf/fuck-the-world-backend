package database

import (
	"fuck-the-world/internal/model"
	"fuck-the-world/unit"
)

func Migrate() {
	err := DB.AutoMigrate(&model.ValidationEmailCode{})
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		unit.HandleError("数据库迁移失败", err)
	}
}
