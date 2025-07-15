package http

import (
	"github.com/wj-stack/job-search/internal/application/service"
	"github.com/wj-stack/job-search/internal/interface/http/handlers"
	"github.com/wj-stack/job-search/internal/interface/http/middleware"
	"github.com/wj-stack/job-search/internal/interface/http/routes"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware 跨域中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

var UserService *service.UserService

// SetupUserService 设置用户服务
func SetupUserService(s *service.UserService) {
	UserService = s
	// 将服务实例传递给 handlers 包
	handlers.UserService = s
}

func SetupRoutes(r *gin.Engine) {
	// 全局添加跨域中间件
	r.Use(CORSMiddleware())
	// 全局添加日志中间件
	r.Use(middleware.LogMiddleware())

	handlers := routes.Handlers{
		LoginHandler:    handlers.LoginHandler,
		RegisterHandler: handlers.RegisterHandler,
		HelloHandler:    handlers.HelloHandler,
	}

	// 设置公开路由
	routes.SetupPublicRoutes(r, handlers)

	// 设置认证路由
	routes.SetupAuthRoutes(r, handlers)
}
