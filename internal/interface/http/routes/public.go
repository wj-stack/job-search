package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupPublicRoutes 设置公开接口路由
func SetupPublicRoutes(r *gin.Engine, handlers Handlers) {
	public := r.Group("/api")
	{
		public.POST("/login", handlers.LoginHandler)
		public.POST("/register", handlers.RegisterHandler)
	}
}
