package repository

import (
	"context"

	"github.com/wj-stack/job-search/internal/domain/model"
)

type JobPaginationQuery struct {
	Page        int    `json:"page"`
	PageSize    int    `json:"page_size"`
	Query       string `json:"query"`
	JobCategory string `json:"job_category"`
	CityInfo    string `json:"city_info"`
}

type JobRepository interface {
	CreateJob(ctx context.Context, title, company string) (*model.Job, error)
	GetJobByID(ctx context.Context, id int) (*model.Job, error)
	ListJobs(ctx context.Context) ([]*model.Job, error)
	UpdateJob(ctx context.Context, id int, title, company string) (*model.Job, error)
	DeleteJob(ctx context.Context, id int) error

	ListJobsWithPagination(ctx context.Context, query JobPaginationQuery) ([]*model.Job, int, error)
}
