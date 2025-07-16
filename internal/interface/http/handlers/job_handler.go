package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wj-stack/job-search/internal/application/service"
)

var JobService *service.JobService

// SearchJobs 岗位搜索接口
func SearchJobs(c *gin.Context) {
	keyword := c.Query("keyword")
	jobType := c.Query("job_type")
	// 此处添加从数据库获取岗位信息的逻辑
	c.JSON(http.StatusOK, gin.H{"message": "实现岗位搜索接口", "keyword": keyword, "jobType": jobType})
}
