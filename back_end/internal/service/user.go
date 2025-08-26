package service

import (
	"context"
	"errors"
	"github.com/tonnyone/go_react_admin/internal/dao"
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

// ...已下沉到dao层...

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
func (s *UserService) Login(ctx context.Context, db *gorm.DB, dto *LoginDTO) error {
	md5Pwd := util.MD5(dto.Password)
	exist, err := s.userDAO.CheckUserByPhoneOrEmailExist(ctx, db, dto.Account, md5Pwd)
	if err != nil {
		return errors.New("内部错误")
	}
	if !exist {
		return errors.New("用户名或密码错误")
	}
	return nil
}
