package routes

import (
	"github.com/wj-stack/job-search/internal/interface/http/handlers"

	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes 设置需要认证的接口路由
func SetupAuthRoutes(r *gin.Engine) {
	public := r.Group("/api")
	{
		public.POST("/login", handlers.LoginHandler)
		public.POST("/register", handlers.RegisterHandler)
	}
}
