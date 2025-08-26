package dao

import (
	"gorm.io/gorm"
)

// Resource 资源表
type Resource struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;size:64"`
	Path      string `gorm:"size:128"`
	Method    string `gorm:"size:16"`
	ParentID  string `gorm:"size:32"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
}

// CreateResource 创建资源
func CreateResource(db *gorm.DB, r *Resource) error {
	return db.Create(r).Error
}

// UpdateResource 更新资源
func UpdateResource(db *gorm.DB, r *Resource) error {
	return db.Save(r).Error
}

// DeleteResourceByID 根据ID删除资源
func DeleteResourceByID(db *gorm.DB, id uint) error {
	return db.Delete(&Resource{}, id).Error
}

// GetResourceByID 根据ID查询资源
func GetResourceByID(db *gorm.DB, id uint) (*Resource, error) {
	var r Resource
	err := db.First(&r, id).Error
	return &r, err
}

// GetResourceByName 根据名称查询资源
func GetResourceByName(db *gorm.DB, name string) (*Resource, error) {
	var r Resource
	err := db.Where("name = ?", name).First(&r).Error
	return &r, err
}

// ListResources 查询所有资源
func ListResources(db *gorm.DB) ([]Resource, error) {
	var resources []Resource
	err := db.Find(&resources).Error
	return resources, err
}

// ListResourcesByParentID 按父ID查询资源
func ListResourcesByParentID(db *gorm.DB, parentID string) ([]Resource, error) {
	var resources []Resource
	err := db.Where("parent_id = ?", parentID).Find(&resources).Error
	return resources, err
}
