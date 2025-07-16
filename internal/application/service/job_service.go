package service

import (
	"github.com/gin-gonic/gin"
	"github.com/wj-stack/job-search/internal/domain/model"
	"github.com/wj-stack/job-search/internal/domain/repository"
)

type JobService struct {
	jobDAO repository.JobRepository
}

func NewJobService(jobDAO repository.JobRepository) *JobService {
	return &JobService{jobDAO: jobDAO}
}

// SearchJobs 实现岗位搜索功能
func (js *JobService) SearchJobs(c *gin.Context, query *repository.JobPaginationQuery) ([]*model.Job, int, error) {
	return js.jobDAO.ListJobsWithPagination(c, query)
}
