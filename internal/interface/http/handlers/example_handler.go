package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/wj-stack/job-search/internal/application/service"
)

// HelloHandler 示例接口处理函数
func HelloHandler(c *gin.Context) {
	result := service.GetHelloMessage()
	c.JSON(200, gin.H{
		"message": result,
	})
	// 保存响应内容到缓存
	c.Set("response", result)
}
