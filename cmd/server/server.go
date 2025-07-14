package main

import (
	"database/sql"
	"fmt"

	"github.com/wj-stack/job-search/internal/application/service"
	"github.com/wj-stack/job-search/internal/infrastructure/dao"
	"github.com/wj-stack/job-search/internal/interface/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	dsn := "your_username:your_password@tcp(127.0.0.1:3306)/your_database"
	var err error
	var db *sql.DB
	if db, err = sql.Open("mysql", dsn); err != nil {
		fmt.Printf("数据库连接初始化失败: %v", err)
		return
	}
	if err = db.Ping(); err != nil {
		fmt.Printf("无法连接到数据库: %v", err)
		return
	}

	// 初始化 DAO 和服务
	userDAO := dao.NewUserDAO(db)
	userService := service.NewUserService(userDAO)
	http.SetupUserService(userService)

	router := gin.Default()
	http.SetupRoutes(router)
	router.Run(":8080")
}
