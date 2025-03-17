package database

import (
	"bookkeeping-server/internal/model"
	"bookkeeping-server/unit"
)

func Migrate() {
	err := DB.AutoMigrate(&model.ValidationEmailCode{})
	if err != nil {
		unit.HandleError("数据库迁移失败", err)
	}
}
