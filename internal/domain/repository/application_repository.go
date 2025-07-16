package repository

import (
	"context"

	"github.com/wj-stack/job-search/internal/domain/model"
)

type ApplicationRepository interface {
	CreateApplication(ctx context.Context, title, company string) (*model.Application, error)
	GetApplicationByID(ctx context.Context, id int) (*model.Application, error)
	ListApplications(ctx context.Context) ([]*model.Application, error)
	DeleteApplication(ctx context.Context, id int) error
}
