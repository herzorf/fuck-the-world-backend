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
	var error error
	for i := 0; i < 5; i++ { // 最多尝试 5 次
		db, err := gorm.Open(mysql.Open(dsnRoot), &gorm.Config{})
		if err == nil {
			DB = db
			log.Println("数据库连接成功")
			break // 连接成功
		}
		error = err
		log.Println(i+1, "数据库连接失败, 重试中...", error)
		time.Sleep(5 * time.Second) // 等待 5 秒再试
	}
	unit.HandleError("数据库连接5次失败", error)
	sqlDB, err := DB.DB()
	if err != nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Minute * 30)
	} else {
		unit.HandleError("获取数据库连接池失败", err)
	}

}
