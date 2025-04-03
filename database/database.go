package database

import (
	"fmt"
	"fuck-the-world/unit"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

func Connect() {
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	dbname := viper.GetString("database.dbname")
	dsnRoot := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	log.Println(dsnRoot)
	var err error
	for i := 0; i < 5; i++ { // 最多尝试 5 次
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break // 连接成功
		}
		log.Println(i+1,"数据库连接失败, 重试中...", err)
		time.Sleep(5 * time.Second) // 等待 5 秒再试
	}
	unit.HandleError("数据库连接失败", err)
	DB = db
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute * 30)
}
