package repository_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/wj-stack/job-search/internal/domain/model"
	"github.com/wj-stack/job-search/internal/domain/repository/mock"
)

func TestCreateUserSuccess(t *testing.T) {
	ctx := context.Background()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mock.NewMockUserRepository(mockCtrl)
	user := &model.User{Username: "test"}
	mockRepo.EXPECT().CreateUser(ctx, user).Return(user, nil)
	createdUser, err := mockRepo.CreateUser(ctx, user)
	assert.NoError(t, err)
	assert.Equal(t, user, createdUser)
}

func TestCreateUserFailure(t *testing.T) {
	ctx := context.Background()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mock.NewMockUserRepository(mockCtrl)
	user := &model.User{Username: "test"}
	createErr := fmt.Errorf("create user failed")
	mockRepo.EXPECT().CreateUser(ctx, user).Return(nil, createErr)
	_, err := mockRepo.CreateUser(ctx, user)
	assert.Error(t, err)
	assert.Equal(t, createErr, err)
}

func TestGetUserByNameSuccess(t *testing.T) {
	ctx := context.Background()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mock.NewMockUserRepository(mockCtrl)
	name := "test"
	user := &model.User{Username: "test"}
	mockRepo.EXPECT().GetUserByName(ctx, name).Return(user, nil)
	_, err := mockRepo.GetUserByName(ctx, name)
	assert.NoError(t, err)
}

func TestGetUserByNameFailure(t *testing.T) {
	ctx := context.Background()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mock.NewMockUserRepository(mockCtrl)
	name := "test"
	getErr := fmt.Errorf("get user by name failed")
	mockRepo.EXPECT().GetUserByName(ctx, name).Return(nil, getErr)
	_, err := mockRepo.GetUserByName(ctx, name)
	assert.Error(t, err)
	assert.Equal(t, getErr, err)
}
