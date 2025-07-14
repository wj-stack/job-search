package service

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/wj-stack/job-search/internal/domain/model"
	"github.com/wj-stack/job-search/internal/infrastructure/dao"
)

// UserService 用户服务
type UserService struct {
	userDAO *dao.UserDAO
}

// NewUserService 创建新的 UserService 实例
func NewUserService(userDAO *dao.UserDAO) *UserService {
	return &UserService{userDAO: userDAO}
}

// Login 用户登录
func (u *UserService) Login(ctx *gin.Context, username, password string) (bool, error) {
	user, err := u.userDAO.GetUserByName(ctx, username)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false, err
	}
	return true, nil
}

// CreateUser 创建新用户
func (u *UserService) CreateUser(ctx *gin.Context, username, password string) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return u.userDAO.CreateUser(ctx, &model.User{
		Username: username,
		Password: string(hashedPassword),
	})
}
