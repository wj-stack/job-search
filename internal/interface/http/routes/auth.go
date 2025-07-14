package routes

import (
	"github.com/wj-stack/job-search/internal/interface/http/middleware"

	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes 设置需要认证的接口路由
func SetupAuthRoutes(r *gin.Engine, handlers Handlers) {
	cache := middleware.NewCache()
	api := r.Group("/api")
	{
		api.Use(middleware.AuthMiddleware())
		api.Use(middleware.CacheMiddleware(cache))
		api.GET("/hello", handlers.HelloHandler)
	}
}
