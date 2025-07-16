package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wj-stack/job-search/internal/interface/http/handlers"
)

// SetupAuthRoutes 设置需要认证的接口路由
func SetupJobRoutes(r *gin.Engine) {
	jobGroup := r.Group("/api/jobs")
	{
		jobGroup.GET("/search", handlers.SearchJobs)
	}
}
