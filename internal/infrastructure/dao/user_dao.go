package dao

import (
	"context"

	"github.com/wj-stack/job-search/ent"
	"github.com/wj-stack/job-search/ent/user"
	"github.com/wj-stack/job-search/internal/domain/model"
)

// UserDAO 定义 User 数据访问对象接口
type UserDAO struct {
	client *ent.Client
}

// NewUserDAO 创建一个新的 UserDAO 实例
func NewUserDAO(client *ent.Client) *UserDAO {
	return &UserDAO{
		client: client,
	}
}

// CreateUser 创建一个新的 User 记录
func (d *UserDAO) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	entUser, err := d.client.User.Create().
		SetUsername(user.Username).
		SetPassword(user.Password).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:       entUser.ID,
		Username: entUser.Username,
		Password: entUser.Password,
	}, nil
}

// GetUserByID 根据 ID 获取 User 记录
func (d *UserDAO) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	entUser, err := d.client.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:       entUser.ID,
		Username: entUser.Username,
		Password: entUser.Password,
	}, nil
}

// GetUserByName 根据用户名获取 User 记录
func (d *UserDAO) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	entUser, err := d.client.User.Query().
		Where(user.Username(name)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:       entUser.ID,
		Username: entUser.Username,
		Password: entUser.Password,
	}, nil
}

// UpdateUser 更新 User 记录
func (d *UserDAO) UpdateUser(ctx context.Context, id int, name string) (*model.User, error) {
	p, err := d.client.User.UpdateOneID(id).
		SetUsername(name).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:       p.ID,
		Username: p.Username,
		Password: p.Password,
	}, nil
}

// DeleteUser 删除 User 记录
func (d *UserDAO) DeleteUser(ctx context.Context, id int) error {
	return d.client.User.DeleteOneID(id).
		Exec(ctx)
}
