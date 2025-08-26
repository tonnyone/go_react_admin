package dao

import (
	"context"
	"gorm.io/gorm"
)

type Role struct {
	ID        string `gorm:"type:char(36);primaryKey"` // uuid主键
	Name      string `gorm:"not null;size:64"`
	Describe  string `gorm:"size:256"`
	Type      string `gorm:"size:32;default:''"` // 角色类型（如 system/custom）
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli"`
	Disabled  bool   `gorm:"default:false"`
}

// RoleDAO 角色数据访问对象
type RoleDAO struct{}

func NewRoleDAO() *RoleDAO {
	return &RoleDAO{}
}

// CreateRole 新增角色
func (dao *RoleDAO) CreateRole(ctx context.Context, db *gorm.DB, role *Role) error {
	return db.WithContext(ctx).Create(role).Error
}

// DeleteRoleByID 根据ID删除角色
func (dao *RoleDAO) DeleteRoleByID(ctx context.Context, db *gorm.DB, id string) error {
	return db.WithContext(ctx).Delete(&Role{}, "id = ?", id).Error
}

// DisableRoleByID 根据ID禁用角色
func (dao *RoleDAO) DisableRoleByID(ctx context.Context, db *gorm.DB, id string) error {
	return db.WithContext(ctx).Model(&Role{}).Where("id = ?", id).Update("disabled", true).Error
}
