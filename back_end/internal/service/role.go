package service

import (
	"context"
	"fmt"

	"github.com/tonnyone/go_react_admin/internal/dao"
	"github.com/tonnyone/go_react_admin/internal/dto"
	"github.com/tonnyone/go_react_admin/internal/util"
	"gorm.io/gorm"
)

// RoleService 角色业务逻辑层
// 负责参数校验、业务处理、调用DAO

type RoleService struct {
	roleDAO *dao.RoleDAO
	db      *gorm.DB
}

func NewRoleService(dao *dao.RoleDAO, db *gorm.DB) *RoleService {
	return &RoleService{roleDAO: dao, db: db}
}

// CreateRole 新增角色
func (s *RoleService) CreateRole(ctx context.Context, roleDTO *dto.RoleDTO) error {

	var role dao.Role
	util.CopyStruct(&role, &roleDTO)
	role.ID = util.GenerateID()
	if role.Type == "" {
		role.Type = string(dto.RoleTypeCustom)
	}
	if s.roleDAO.CheckRoleExist(ctx, s.db, role.Name) {
		return fmt.Errorf("角色已存在: %s", role.Name)
	}
	// 可加业务校验
	return s.roleDAO.CreateRole(ctx, s.db, &role)
}

// UpdateRole 修改角色
func (s *RoleService) UpdateRole(ctx context.Context, id string, req *dao.Role) error {
	var role dao.Role
	if err := s.db.WithContext(ctx).First(&role, "id = ?", id).Error; err != nil {
		return err
	}
	role.Name = req.Name
	role.Describe = req.Describe
	role.Type = req.Type
	return s.db.WithContext(ctx).Save(&role).Error
}

// DeleteRole 删除角色
func (s *RoleService) DeleteRole(ctx context.Context, id string) error {
	return s.roleDAO.DeleteRoleByID(ctx, s.db, id)
}

// ListRoles 分页查询角色
func (s *RoleService) ListRoles(ctx context.Context, page, pageSize int) ([]dao.Role, int64, error) {
	var roles []dao.Role
	var total int64
	s.db.WithContext(ctx).Model(&dao.Role{}).Count(&total)
	s.db.WithContext(ctx).Offset((page - 1) * pageSize).Limit(pageSize).Find(&roles)
	return roles, total, nil
}
