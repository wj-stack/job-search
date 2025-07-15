package repository

import (
	"context"

	"github.com/wj-stack/job-search/internal/domain/model"
)

// UserRepository 用户仓库接口
//
//go:generate mockgen --source ./user_repository.go --destination ./mock/mock_user_repository.go --package mock
type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}
