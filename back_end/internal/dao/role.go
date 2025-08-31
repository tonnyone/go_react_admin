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

// FindRolesByUserIDs 根据一组用户ID，查询他们所有的角色
func (d *RoleDAO) FindRolesByUserIDs(ctx context.Context, db *gorm.DB, userIDs []string) (map[string][]Role, error) {
	// 定义一个临时的结构体来接收 JOIN 查询的结果
	type UserRoleResult struct {
		UserID   string
		RoleID   string
		RoleName string
	}
	var results []UserRoleResult
	// 执行原生 SQL JOIN 查询
	// 这里我们直接操作 user_roles 表
	err := db.WithContext(ctx).Table("roles as r").
		Select("ur.user_id, r.id as role_id, r.name as role_name").
		Joins("JOIN user_roles as ur ON r.id = ur.role_id").
		Where("ur.user_id IN ?", userIDs).
		Scan(&results).Error

	if err != nil {
		return nil, err
	}
	// 将扁平化的查询结果组装成 map
	rolesMap := make(map[string][]Role)
	for _, r := range results {
		role := Role{
			ID:   r.RoleID,
			Name: r.RoleName,
		}
		rolesMap[r.UserID] = append(rolesMap[r.UserID], role)
	}
	return rolesMap, nil
}
