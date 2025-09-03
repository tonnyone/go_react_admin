package dao

import (
	"gorm.io/gorm"
)

// CreateRoleResource 创建角色-资源关联
func CreateRoleResource(db *gorm.DB, rr *RoleResource) error {
	return db.Create(rr).Error
}

// DeleteRoleResourceByID 根据ID删除关联
func DeleteRoleResourceByID(db *gorm.DB, id uint) error {
	return db.Delete(&RoleResource{}, id).Error
}

// DeleteRoleResource 删除指定角色与资源的关联
func DeleteRoleResource(db *gorm.DB, roleID, resourceID uint) error {
	return db.Where("role_id = ? AND resource_id = ?", roleID, resourceID).Delete(&RoleResource{}).Error
}

// GetRoleResourcesByRoleID 查询某角色的所有资源ID
func GetRoleResourcesByRoleID(db *gorm.DB, roleID uint) ([]uint, error) {
	var resourceIDs []uint
	err := db.Model(&RoleResource{}).Where("role_id = ?", roleID).Pluck("resource_id", &resourceIDs).Error
	return resourceIDs, err
}

// GetRoleResourcesByResourceID 查询某资源的所有角色ID
func GetRoleResourcesByResourceID(db *gorm.DB, resourceID uint) ([]uint, error) {
	var roleIDs []uint
	err := db.Model(&RoleResource{}).Where("resource_id = ?", resourceID).Pluck("role_id", &roleIDs).Error
	return roleIDs, err
}

// BatchCreateRoleResources 批量创建角色-资源关联
func BatchCreateRoleResources(db *gorm.DB, rrs []RoleResource) error {
	return db.Create(&rrs).Error
}

// DeleteRoleResourcesByRoleID 删除某角色的所有资源关联
func DeleteRoleResourcesByRoleID(db *gorm.DB, roleID uint) error {
	return db.Where("role_id = ?", roleID).Delete(&RoleResource{}).Error
}

// DeleteRoleResourcesByResourceID 删除某资源的所有角色关联
func DeleteRoleResourcesByResourceID(db *gorm.DB, resourceID uint) error {
	return db.Where("resource_id = ?", resourceID).Delete(&RoleResource{}).Error
}
