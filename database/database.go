package database

import (
	"bookkeeping-server/unit"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	dbname := viper.GetString("database.dbname")
	dsnRoot := fmt.Sprintf("%s:%s@tcp(%s:%d)/", username, password, host, port)
	db, err := gorm.Open(mysql.Open(dsnRoot), &gorm.Config{})
	unit.HandleError("数据库连接失败", err)
	DB = db
	// 如果数据库不存在就创建数据库
	DB.Exec("CREATE DATABASE IF NOT EXISTS " + dbname + " CHARSET utf8mb4 COLLATE utf8mb4_general_ci;")
	if err := DB.Exec(fmt.Sprintf("USE %s", dbname)).Error; err != nil {
		unit.HandleError("选择数据库失败", err)
	}
}
