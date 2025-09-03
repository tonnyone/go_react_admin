package dao

type User struct {
	ID         string `gorm:"type:char(22);primaryKey"` // uuid主键
	Username   string `gorm:"not null;size:64"`
	Email      string `gorm:"unique;size:128"`
	Phone      string `gorm:"unique;size:64"`
	Password   string `gorm:"not null;size:128" json:"password,omitempty"` // 忽略掉
	Department string `gorm:"size:128"`
	CreatedAt  int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt  int64  `gorm:"autoUpdateTime:milli"`
	Disabled   bool   `gorm:"default:false"`
}

func (r *User) TableName() string {
	return "sys_user"
}

type Role struct {
	ID        string `gorm:"type:char(22);primaryKey"` // uuid主键
	Name      string `gorm:"unique;size:64"`
	Describe  string `gorm:"size:256"`
	Type      string `gorm:"size:32;default:''"` // 角色类型（如 system/custom）
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli"`
	Disabled  bool   `gorm:"default:false"`
}

func (r *Role) TableName() string {
	return "sys_role"
}

type UserRole struct {
	UserID    string `gorm:"type:char(22);not null;primaryKey"`
	RoleID    string `gorm:"type:char(22);not null;primaryKey"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	CreatedBy string `gorm:"type:char(22);not null;"`
}

func (r *UserRole) TableName() string {
	return "sys_user_role"
}

// RoleMenu 角色-菜单关联表
type RoleMenu struct {
	ID        string `gorm:"type:char(22);primaryKey"`
	RoleID    string `gorm:"type:char(22);not null;index"`
	MenuID    string `gorm:"type:char(22);not null;index"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
}

func (r *RoleMenu) TableName() string {
	return "sys_role_menu"
}

type RoleResource struct {
	ID         string `gorm:"type:char(22);primaryKey" json:"id"`
	RoleID     string `gorm:"type:char(22);not null;index;comment:角色ID" json:"role_id"`
	ResourceID string `gorm:"type:char(22);not null;index;comment:资源ID" json:"resource_id"`
	CreatedAt  int64  `gorm:"autoCreateTime:milli"`
}

// TableName 自定义表名
func (RoleResource) TableName() string {
	return "role_resource"
}

type Menu struct {
	ID        string `gorm:"type:char(22);primaryKey"`
	Name      string `gorm:"not null;size:64"`
	ParentID  string `gorm:"type:char(22);"`
	Path      string `gorm:"size:128"` // 前端路由
	Type      string `gorm:"size:16"`  // 菜单/按钮
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
}

// Resource 资源表
type Resource struct {
	ID        string `gorm:"type:char(22);primaryKey"`
	Name      string `gorm:"not null;size:64"`
	Path      string `gorm:"size:128"`
	Method    string `gorm:"size:16"`
	ParentID  string `gorm:"type:char(22);"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
}
