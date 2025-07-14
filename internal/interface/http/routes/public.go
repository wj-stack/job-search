package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupPublicRoutes 设置公开接口路由
func SetupPublicRoutes(r *gin.Engine, loginHandler, registerHandler gin.HandlerFunc) {
	public := r.Group("/api")
	{
		public.POST("/login", loginHandler)
		public.POST("/register", registerHandler)
	}
}
