package dao

import (
	"context"
	"gorm.io/gorm"
)

// RoleMenuDAO 角色-菜单关联表操作
type RoleMenuDAO struct{}

func NewRoleMenuDAO() *RoleMenuDAO {
	return &RoleMenuDAO{}
}

// AddRoleMenu 新增角色-菜单关联
func (dao *RoleMenuDAO) AddRoleMenu(ctx context.Context, db *gorm.DB, roleID, menuID, id string) error {
	rm := RoleMenu{ID: id, RoleID: roleID, MenuID: menuID}
	return db.WithContext(ctx).Create(&rm).Error
}

// DeleteRoleMenu 删除角色-菜单关联
func (dao *RoleMenuDAO) DeleteRoleMenu(ctx context.Context, db *gorm.DB, roleID, menuID string) error {
	return db.WithContext(ctx).Where("role_id = ? AND menu_id = ?", roleID, menuID).Delete(&RoleMenu{}).Error
}

// ListMenusByRole 查询角色所有菜单ID
func (dao *RoleMenuDAO) ListMenusByRole(ctx context.Context, db *gorm.DB, roleID string) ([]string, error) {
	var rms []RoleMenu
	if err := db.WithContext(ctx).Where("role_id = ?", roleID).Find(&rms).Error; err != nil {
		return nil, err
	}
	ids := make([]string, 0, len(rms))
	for _, rm := range rms {
		ids = append(ids, rm.MenuID)
	}
	return ids, nil
}
