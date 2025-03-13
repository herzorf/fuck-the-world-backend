package database

import (
	"bookkeeping-server/unit"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
	dbRoot.Exec("CREATE DATABASE IF NOT EXISTS " + dbname + " CHARSET utf8mb4 COLLATE utf8mb4_general_ci;")
}
