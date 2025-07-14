package dao

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// InitDB 初始化数据库连接
func InitDB(dsn string) {
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("数据库连接初始化失败: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}
	log.Println("数据库连接成功")
}
