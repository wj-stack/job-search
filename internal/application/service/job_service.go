package service

import (
	"github.com/gin-gonic/gin"
	"github.com/wj-stack/job-search/internal/domain/repository"
)

type JobService struct {
	jobDAO repository.JobRepository
}

func NewJobService(jobDAO repository.JobRepository) *JobService {
	return &JobService{jobDAO: jobDAO}
}

// SearchJobs 实现 Web 上的岗位搜索功能
func (js *JobService) SearchJobs(c *gin.Context) {
	keyword := c.Query("keyword")
	jobType := c.Query("job_type")
	// 构建查询条件，从数据库获取岗位信息
	// 示例返回，实际需替换为真实数据查询逻辑
	result := []map[string]interface{}{}
	// 模拟查询逻辑
	if keyword != "" || jobType != "" {
		result = append(result, map[string]interface{}{
			"id":      1,
			"title":   "后端开发工程师",
			"company": "XX 公司",
		})
	}
	c.JSON(200, result)
}
