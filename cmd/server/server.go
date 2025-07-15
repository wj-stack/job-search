package main

import (
	"flag"
	"fmt"

	"github.com/wj-stack/job-search/internal/application/service"
	"github.com/wj-stack/job-search/internal/infrastructure/config"
	"github.com/wj-stack/job-search/internal/infrastructure/dao"
	"github.com/wj-stack/job-search/internal/interface/http"

	"github.com/gin-gonic/gin"
)

func main() {
	configPath := flag.String("config", "../../config.yaml", "Path to the configuration file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		return
	}

	client := dao.InitDB(cfg.Database.Driver, cfg.Database.Source)
	defer client.Close()
	// 初始化 DAO 和服务
	userDAO := dao.NewUserDAO(client)
	userService := service.NewUserService(userDAO)
	http.SetupUserService(userService)

	router := gin.Default()
	http.SetupRoutes(router)
	router.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}
