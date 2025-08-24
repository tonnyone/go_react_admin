package dao

// 用户表
// User 结构体
// 用户和角色多对多
// 用户和菜单多对多
// 用户和资源多对多

// type User struct {
// 	ID        uint       `gorm:"primaryKey"`
// 	Username  string     `gorm:"unique;not null;size:64"`
// 	Password  string     `gorm:"not null;size:128"`
// 	Email     string     `gorm:"size:128"`
// 	Roles     []Role     `gorm:"many2many:user_roles;"`
// 	Menus     []Menu     `gorm:"many2many:user_menus;"`
// 	Resources []Resource `gorm:"many2many:user_resources;"`
// 	CreatedAt gorm.DeletedAt
// }

// type Role struct {
// 	ID        uint       `gorm:"primaryKey"`
// 	Name      string     `gorm:"unique;not null;size:64"`
// 	Users     []User     `gorm:"many2many:user_roles;"`
// 	Menus     []Menu     `gorm:"many2many:role_menus;"`
// 	Resources []Resource `gorm:"many2many:role_resources;"`
// }

// type Menu struct {
// 	ID       uint   `gorm:"primaryKey"`
// 	Name     string `gorm:"not null;size:64"`
// 	ParentID *uint
// 	Children []Menu `gorm:"foreignKey:ParentID"`
// 	Path     string `gorm:"size:128"`
// 	Type     string `gorm:"size:16"` // 菜单/按钮
// 	Roles    []Role `gorm:"many2many:role_menus;"`
// 	Users    []User `gorm:"many2many:user_menus;"`
// }

// type Resource struct {
// 	ID     uint   `gorm:"primaryKey"`
// 	Name   string `gorm:"not null;size:64"`
// 	Path   string `gorm:"size:128"`
// 	Method string `gorm:"size:16"`
// 	Roles  []Role `gorm:"many2many:role_resources;"`
// 	Users  []User `gorm:"many2many:user_resources;"`
// }
