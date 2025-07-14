package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/wj-stack/job-search/internal/application/service"
)

var UserService *service.UserService

// LoginHandler 用户登录处理函数
func LoginHandler(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "请求参数错误"})
		return
	}

	loggedIn, err := UserService.Login(credentials.Username, credentials.Password)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "服务器错误"})
		return
	}

	if loggedIn {
		c.JSON(200, gin.H{"message": "登录成功"})
	} else {
		c.JSON(401, gin.H{"error": "用户名或密码错误"})
	}
}

// RegisterHandler 用户注册处理函数
func RegisterHandler(c *gin.Context) {
	var userData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "请求参数错误"})
		return
	}

	if err := UserService.CreateUser(userData.Username, userData.Password); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "注册失败"})
		return
	}

	c.JSON(200, gin.H{"message": "注册成功"})
}
