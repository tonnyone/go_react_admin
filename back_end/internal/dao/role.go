package dao

import (
	"context"
	"gorm.io/gorm"
)

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
		Role
		UserID string
	}
	var results []UserRoleResult
	// 执行原生 SQL JOIN 查询
	// 这里我们直接操作 user_roles 表
	err := db.WithContext(ctx).Table("sys_user_role as ur").
		Select("ur.user_id as user_id, ur.role_id as role_id, r.name as name, r.id as id").
		Joins("JOIN sys_role as r ON r.id = ur.role_id").
		Where("ur.user_id IN ?", userIDs).
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	// 将扁平化的查询结果组装成 map
	rolesMap := make(map[string][]Role)
	for _, r := range results {
		rolesMap[r.UserID] = append(rolesMap[r.UserID], r.Role)
	}
	return rolesMap, nil
}

// CheckRoleExist 检查角色名是否已存在
func (dao *RoleDAO) CheckRoleExist(ctx context.Context, db *gorm.DB, name string) bool {
	var count int64
	db.WithContext(ctx).Model(&Role{}).Where("name = ?", name).Count(&count)
	return count > 0
}
