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

// SetupUserService 设置用户服务
func SetupUserService(s *service.UserService) {
	// 将服务实例传递给 handlers 包
	handlers.UserService = s
}

// SetupJobService 设置Job服务
func SetupJobService(s *service.JobService) {
	// 将服务实例传递给 handlers 包
	handlers.JobService = s
}

func SetupRoutes(r *gin.Engine) {
	// 全局添加跨域中间件
	r.Use(CORSMiddleware())
	// 全局添加日志中间件
	r.Use(middleware.LogMiddleware())

	// 设置公开路由
	routes.SetupPublicRoutes(r)

	// 设置认证路由
	routes.SetupAuthRoutes(r)

	// 注册岗位路由
	routes.SetupJobRoutes(r)

	// 注册投递路由
	routes.SetupApplicationRoutes(r)
}
