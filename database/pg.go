package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var (
	host     = "localhost"
	port     = 5432
	username = "mangosteen"
	password = "123456"
	dbname   = "mangosteen_dev"
)

func Connect() {
	// 创建连接字符串
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", username, password, host, port, dbname)

	// 打开数据库连接
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening connection to the database: ", err)
	}
	DB = db
	defer Close()

	// 测试连接
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// 创建表
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100)
	);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Error creating table: ", err)
	}

	// 插入数据
	insertQuery := `INSERT INTO users (name) VALUES ($1), ($2), ($3)`
	_, err = db.Exec(insertQuery, "Alice", "Bob", "Charlie")
	if err != nil {
		log.Fatal("Error inserting data: ", err)
	}

	// 查询数据
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal("Error executing query: ", err)
	}
	defer rows.Close()

	// 打印查询结果
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal("Error scanning row: ", err)
		}
		fmt.Println(id, name)
	}

	// 检查迭代过程中的错误
	err = rows.Err()
	if err != nil {
		log.Fatal("Error after iterating rows: ", err)
	}
}

func Close() {
	DB.Close()
}
