package repository

import (
	"context"

	"github.com/wj-stack/job-search/internal/domain/model"
)

type JobRepository interface {
	CreateJob(ctx context.Context, title, company string) (*model.Job, error)
	GetJobByID(ctx context.Context, id int) (*model.Job, error)
	ListJobs(ctx context.Context) ([]*model.Job, error)
	UpdateJob(ctx context.Context, id int, title, company string) (*model.Job, error)
	DeleteJob(ctx context.Context, id int) error
}
