package service

import (
	"context"
	"errors"
	"github.com/tonnyone/go_react_admin/internal/dao"

	"gorm.io/gorm"
)

type UserService struct {
	userDAO *dao.UserDAO
}

// NewUserService 构造函数，注入 UserDAO
func NewUserService(userDAO *dao.UserDAO) *UserService {
	return &UserService{userDAO: userDAO}
}

// LoginDTO 定义登录业务参数
type LoginDTO struct {
	Username string
	Password string
}

// Login 处理登录逻辑, 返回 token
func (s *UserService) Login(ctx context.Context, db *gorm.DB, dto *LoginDTO) (string, error) {
	exist, err := s.userDAO.CheckUserExist(ctx, db, dto.Username, dto.Password)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", errors.New("用户名或密码错误")
	}
	// TODO: 生成token
	return "mock-token", nil
}
