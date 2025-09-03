package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/samber/lo"
	"github.com/tonnyone/go_react_admin/internal/dao"
	"github.com/tonnyone/go_react_admin/internal/dto"
	"github.com/tonnyone/go_react_admin/internal/util"
	"gorm.io/gorm"
)

type UserService struct {
	db      *gorm.DB
	userDAO *dao.UserDAO
	roleDAO *dao.RoleDAO
}

// NewUserService 构造函数，注入 UserDAO
func NewUserService(userDAO *dao.UserDAO, db *gorm.DB) *UserService {
	return &UserService{userDAO: userDAO, db: db}
}

// LoginDTO 定义登录业务参数
type LoginDTO struct {
	Account  string
	Password string
}

// 注册用户
func (s *UserService) Register(ctx context.Context, phone, email, password string) error {
	if phone != "" {
		exist, err := s.userDAO.CheckPhoneExist(ctx, s.db, phone)
		if err != nil {
			return errors.New("内部错误")
		}
		if exist {
			return errors.New("手机号已存在")
		}
	}
	if email != "" {
		exist, err := s.userDAO.CheckEmailExist(ctx, s.db, email)
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
	if err := s.db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// Login 处理登录逻辑, 返回 token
func (s *UserService) Login(ctx context.Context, dto *LoginDTO) (string, error) {
	md5Pwd := util.MD5(dto.Password)
	exist, err := s.userDAO.CheckUserByPhoneOrEmailExist(ctx, s.db, dto.Account, md5Pwd)
	if err != nil {
		return "", errors.New("内部错误")
	}
	if !exist {
		return "", errors.New("用户名或密码错误")
	}
	return util.GenAuthToken(dto.Account, dto.Password), nil
}

// GetUsers 获取用户列表
func (s *UserService) GetUsers(ctx context.Context, pager *dto.Pager) ([]dto.UserDTO, int64, error) {

	users, total, err := s.userDAO.GetUsers(ctx, s.db, pager)
	if err != nil {
		return nil, 0, fmt.Errorf("GetUsers failed: %w", err)
	}
	userIDs := lo.Map(users, func(user dao.User, _ int) string {
		return user.ID
	})
	rolesMap, err := s.roleDAO.FindRolesByUserIDs(ctx, s.db, userIDs)
	if err != nil {
		return nil, 0, fmt.Errorf("FindRolesByUserIDs failed: %w", err)
	}
	userDTOs := make([]dto.UserDTO, 0, len(users))
	for _, user := range users {
		var userDTO dto.UserDTO
		if err := util.CopyStruct(&userDTO, &user); err != nil {
			return nil, 0, fmt.Errorf("CopyStruct user to userDTO failed: %w", err)
		}
		if roles, ok := rolesMap[user.ID]; ok {
			if err := util.CopyStruct(&userDTO.Roles, &roles); err != nil {
				return nil, 0, fmt.Errorf("CopyStruct roles to userDTO.Roles failed: %w", err)
			}
		}
		userDTOs = append(userDTOs, userDTO)
	}
	return userDTOs, total, nil
}

// BindRolesToUser handles the business logic of binding roles to a user.
func (s *UserService) BindRolesToUser(ctx context.Context, req *dto.BindUserRolesReq) error {
	if req.UserID == "" || len(req.RoleIDs) <= 0 {
		return fmt.Errorf("userId 或者 roleIDs 为空")
	}
	return s.db.Transaction(func(tx *gorm.DB) error {
		user, err := s.userDAO.GetUser(ctx, tx, req.UserID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errors.New("user not found")
			}
			return err
		}
		return s.userDAO.AppendUserRoles(ctx, tx, user.ID, req.RoleIDs)
	})
}
