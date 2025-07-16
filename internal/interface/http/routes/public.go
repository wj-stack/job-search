package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wj-stack/job-search/internal/interface/http/handlers"
)

// SetupPublicRoutes 设置公开接口路由
func SetupPublicRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/hello", handlers.HelloHandler)
	}
}
