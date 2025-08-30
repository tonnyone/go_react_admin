package service

import (
	"context"
	"errors"
	"github.com/tonnyone/go_react_admin/internal/dao"
	"github.com/tonnyone/go_react_admin/internal/dto"
	"github.com/tonnyone/go_react_admin/internal/util"
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
	Account  string
	Password string
}

// UserDTO 定义用户业务参数
type UserDTO struct {
	ID         string     `json:"id,omitempty"` // uuid主键
	Username   string     `json:"username,omitempty"`
	Email      string     `json:"email,omitempty"`
	Phone      string     `json:"phone,omitempty"`
	Password   string     `json:"password,omitempty"` // 忽略掉
	Department string     `json:"department,omitempty"`
	CreatedAt  int64      `json:"created_at,omitempty"`
	UpdatedAt  int64      `json:"updated_at,omitempty"`
	Disabled   bool       `json:"disabled,omitempty"`
	Roles      []dao.Role `json:"roles,omitempty"` // 关联的角色
}

type UserRoleDTO struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// 注册用户
func (s *UserService) Register(ctx context.Context, db *gorm.DB, phone, email, password string) error {
	if phone != "" {
		exist, err := s.userDAO.CheckPhoneExist(ctx, db, phone)
		if err != nil {
			return errors.New("内部错误")
		}
		if exist {
			return errors.New("手机号已存在")
		}
	}
	if email != "" {
		exist, err := s.userDAO.CheckEmailExist(ctx, db, email)
		if err != nil {
			return errors.New("内部错误")
		}
		if exist {
			return errors.New("邮箱已存在")
		}
	}
	user := dao.User{
		ID:       util.GenerateID(),
		Username: "无用户名",
		Phone:    phone,
		Email:    email,
		Password: util.MD5(password),
	}
	if err := db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// Login 处理登录逻辑, 返回 token
func (s *UserService) Login(ctx context.Context, db *gorm.DB, dto *LoginDTO) (string, error) {
	md5Pwd := util.MD5(dto.Password)
	exist, err := s.userDAO.CheckUserByPhoneOrEmailExist(ctx, db, dto.Account, md5Pwd)
	if err != nil {
		return "", errors.New("内部错误")
	}
	if !exist {
		return "", errors.New("用户名或密码错误")
	}
	return util.GenAuthToken(dto.Account, dto.Password), nil
}

// GetUsers 获取用户列表
func (s *UserService) GetUsers(ctx context.Context, db *gorm.DB, pager *dto.Pager) ([]dao.User, int64, error) {
	return s.userDAO.GetUsers(ctx, db, pager)
}
