package service

import (
	"github.com/wj-stack/job-search/internal/domain/model"
	"github.com/wj-stack/job-search/internal/infrastructure/dao"
	"golang.org/x/crypto/bcrypt"
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
func (u *UserService) Login(username, password string) (bool, error) {
	user, err := u.userDAO.GetUserByUsername(username)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, nil
	}
	return true, nil
}

// CreateUser 创建新用户
func (u *UserService) CreateUser(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &model.User{
		Username: username,
		Password: string(hashedPassword),
	}
	return u.userDAO.CreateUser(user)
}
