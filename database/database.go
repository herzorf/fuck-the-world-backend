package database

import (
	"bookkeeping-server/unit"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type ValidationCode struct {
	ID        uint       `gorm:"primaryKey"`
	Code      string     `gorm:"size:20;not null"`
	Email     string     `gorm:"size:255;not null"`
	UsedAt    *time.Time // 可以为空，代表未使用
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}

var (
	host     = "localhost"
	port     = 3306
	username = "root"
	password = "123456"
	dbname   = "bookKeeping_db_dev"
)
var DB *gorm.DB

func Connect() {
	dsnRoot := fmt.Sprintf("%s:%s@tcp(%s:%d)/", username, password, host, port)
	dbRoot, err := gorm.Open(mysql.Open(dsnRoot), &gorm.Config{})
	unit.HandleError("数据库连接失败", err)
	DB = dbRoot
	// 如果数据库不存在就创建数据库
	DB.Exec("CREATE DATABASE IF NOT EXISTS " + dbname + " CHARSET utf8mb4 COLLATE utf8mb4_general_ci;")
	DB.Exec(fmt.Sprintf("USE %s", dbname))
}

func Migrate() {
	DB.AutoMigrate(&ValidationCode{})
}
