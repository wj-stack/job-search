package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wj-stack/job-search/internal/interface/http/handlers"
)

// SetupApplicationRoutes 设置需要认证的接口路由
func SetupApplicationRoutes(r *gin.Engine) {
	appGroup := r.Group("/api/applications")
	{
		appGroup.POST("", handlers.CreateApplication)
		appGroup.GET("", handlers.GetApplications)
	}

	resumeGroup := r.Group("/api/resume")
	{
		resumeGroup.POST("/upload", handlers.UploadResume)
	}
}
