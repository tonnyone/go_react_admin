package dao

import (
	"gorm.io/gorm"
)

// CreateMenu 创建菜单
func CreateMenu(db *gorm.DB, m *Menu) error {
	return db.Create(m).Error
}

// UpdateMenu 更新菜单
func UpdateMenu(db *gorm.DB, m *Menu) error {
	return db.Save(m).Error
}

// DeleteMenuByID 根据ID删除菜单
func DeleteMenuByID(db *gorm.DB, id string) error {
	return db.Delete(&Menu{}, id).Error
}

// GetMenuByID 根据ID查询菜单
func GetMenuByID(db *gorm.DB, id string) (*Menu, error) {
	var m Menu
	err := db.First(&m, id).Error
	return &m, err
}

// GetMenuByName 根据名称查询菜单
func GetMenuByName(db *gorm.DB, name string) (*Menu, error) {
	var m Menu
	err := db.Where("name = ?", name).First(&m).Error
	return &m, err
}

// ListMenus 查询所有菜单
func ListMenus(db *gorm.DB) ([]Menu, error) {
	var menus []Menu
	err := db.Find(&menus).Error
	return menus, err
}

// ListMenusByParentID 按父ID查询菜单
func ListMenusByParentID(db *gorm.DB, parentID string) ([]Menu, error) {
	var menus []Menu
	err := db.Where("parent_id = ?", parentID).Find(&menus).Error
	return menus, err
}

// ListMenusByType 按类型查询菜单
func ListMenusByType(db *gorm.DB, typ string) ([]Menu, error) {
	var menus []Menu
	err := db.Where("type = ?", typ).Find(&menus).Error
	return menus, err
}
