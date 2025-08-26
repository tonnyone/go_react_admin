package service

import (
	"context"
	"github.com/tonnyone/go_react_admin/internal/dao"
	"gorm.io/gorm"
)

// RoleService 角色业务逻辑层
// 负责参数校验、业务处理、调用DAO

type RoleService struct {
	roleDAO *dao.RoleDAO
}

func NewRoleService() *RoleService {
	return &RoleService{roleDAO: dao.NewRoleDAO()}
}

// CreateRole 新增角色
func (s *RoleService) CreateRole(ctx context.Context, db *gorm.DB, role *dao.Role) error {

	// 可加业务校验
	return s.roleDAO.CreateRole(ctx, db, role)
}

// UpdateRole 修改角色
func (s *RoleService) UpdateRole(ctx context.Context, db *gorm.DB, id string, req *dao.Role) error {
	var role dao.Role
	if err := db.WithContext(ctx).First(&role, "id = ?", id).Error; err != nil {
		return err
	}
	role.Name = req.Name
	role.Describe = req.Describe
	role.Type = req.Type
	return db.WithContext(ctx).Save(&role).Error
}

// DeleteRole 删除角色
func (s *RoleService) DeleteRole(ctx context.Context, db *gorm.DB, id string) error {
	return s.roleDAO.DeleteRoleByID(ctx, db, id)
}

// ListRoles 分页查询角色
func (s *RoleService) ListRoles(ctx context.Context, db *gorm.DB, page, pageSize int) ([]dao.Role, int64, error) {
	var roles []dao.Role
	var total int64
	db.WithContext(ctx).Model(&dao.Role{}).Count(&total)
	db.WithContext(ctx).Offset((page - 1) * pageSize).Limit(pageSize).Find(&roles)
	return roles, total, nil
}
