package database

import (
	"fuck-the-world/internal/model"
	"fuck-the-world/utils"
)

func Migrate() {
	err := DB.AutoMigrate(&model.ValidationEmailCode{})
	err = DB.AutoMigrate(&model.User{})
	err = DB.AutoMigrate(&model.Map{})
	err = DB.AutoMigrate(&model.Player{})
	err = DB.AutoMigrate(&model.MapPlayerSave{})
	if err != nil {
		utils.HandleError("数据库迁移失败", err)
	}
}
