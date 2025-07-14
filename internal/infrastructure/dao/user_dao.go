package dao

import (
	"database/sql"
	"log"

	"github.com/wj-stack/job-search/internal/domain/model"
)

// UserDAO 用户数据访问对象
type UserDAO struct {
	DB *sql.DB
}

// NewUserDAO 创建新的 UserDAO 实例
func NewUserDAO(db *sql.DB) *UserDAO {
	return &UserDAO{DB: db}
}

// GetUserByUsername 根据用户名获取用户信息
func (u *UserDAO) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := u.DB.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("查询用户失败: %v", err)
		return nil, err
	}
	return &user, nil
}

// CreateUser 创建新用户
func (u *UserDAO) CreateUser(user *model.User) error {
	_, err := u.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		log.Printf("创建用户失败: %v", err)
		return err
	}
	return nil
}
