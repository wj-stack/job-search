package middleware

import "github.com/gin-gonic/gin"

// AuthMiddleware 认证中间件示例
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里添加认证逻辑，示例仅作演示
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "未提供认证信息"})
			return
		}
		c.Next()
	}
}
