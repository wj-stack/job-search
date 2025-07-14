package routes

import "github.com/gin-gonic/gin"

// Handlers 保存所有的 HandlerFunc
type Handlers struct {
	LoginHandler    gin.HandlerFunc
	RegisterHandler gin.HandlerFunc
	HelloHandler    gin.HandlerFunc
}