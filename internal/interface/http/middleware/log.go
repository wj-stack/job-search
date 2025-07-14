package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

// LogMiddleware 日志中间件示例
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		// 记录请求日志，示例仅作演示
		logEntry := map[string]interface{}{
			"method":  c.Request.Method,
			"path":    c.Request.URL.Path,
			"status":  c.Writer.Status(),
			"latency": latency.String(),
		}
		// 实际项目中可以将日志写入文件或发送到日志服务
		_ = logEntry
	}
}
