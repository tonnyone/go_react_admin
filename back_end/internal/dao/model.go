package dao

type User struct {
	ID         string `gorm:"type:char(22);primaryKey" json:"id"` // uuid主键
	Username   string `gorm:"not null;size:64" json:"username"`
	Email      string `gorm:"unique;size:128" json:"email"`
	Phone      string `gorm:"unique;size:64" json:"phone"`
	Password   string `gorm:"not null;size:128" json:"password,omitempty"` // 忽略掉
	Department string `gorm:"size:128" json:"department"`
	CreatedAt  int64  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt  int64  `gorm:"autoUpdateTime:milli" json:"updated_at"`
	Disabled   bool   `gorm:"default:false" json:"disabled"`
}

func (r *User) TableName() string {
	return "sys_user"
}

type Role struct {
	ID        string `gorm:"type:char(22);primaryKey" json:"id"` // uuid主键
	Name      string `gorm:"unique;size:64" json:"name"`
	Describe  string `gorm:"size:256" json:"describe"`
	Type      string `gorm:"size:32;default:''" json:"type"` // 角色类型（如 system/custom）
	CreatedAt int64  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli" json:"updated_at"`
	Disabled  bool   `gorm:"default:false" json:"disabled"`
}

func (r *Role) TableName() string {
	return "sys_role"
}

type UserRole struct {
	UserID    string `gorm:"type:char(22);not null;primaryKey" json:"user_id"`
	RoleID    string `gorm:"type:char(22);not null;primaryKey" json:"role_id"`
	CreatedAt int64  `gorm:"autoCreateTime:milli" json:"created_at"`
	CreatedBy string `gorm:"type:char(22);not null;" json:"created_by"`
}

func (r *UserRole) TableName() string {
	return "sys_user_role"
}

// RoleMenu 角色-菜单关联表
type RoleMenu struct {
	ID        string `gorm:"type:char(22);primaryKey" json:"id"`
	RoleID    string `gorm:"type:char(22);not null;index" json:"role_id"`
	MenuID    string `gorm:"type:char(22);not null;index" json:"menu_id"`
	CreatedAt int64  `gorm:"autoCreateTime:milli" json:"created_at"`
}

func (r *RoleMenu) TableName() string {
	return "sys_role_menu"
}

type RoleResource struct {
	ID         string `gorm:"type:char(22);primaryKey" json:"id"`
	RoleID     string `gorm:"type:char(22);not null;index;comment:角色ID" json:"role_id"`
	ResourceID string `gorm:"type:char(22);not null;index;comment:资源ID" json:"resource_id"`
	CreatedAt  int64  `gorm:"autoCreateTime:milli" json:"created_at"`
}

// TableName 自定义表名
func (RoleResource) TableName() string {
	return "sys_role_resource"
}

type Menu struct {
	ID        string `gorm:"type:char(22);primaryKey" json:"id"`
	Name      string `gorm:"not null;size:64" json:"name"`
	ParentID  string `gorm:"type:char(22);" json:"parent_id"`
	Path      string `gorm:"size:128" json:"path"` // 前端路由
	Type      string `gorm:"size:16" json:"type"`  // 菜单/按钮
	CreatedAt int64  `gorm:"autoCreateTime:milli" json:"created_at"`
}

// TableName 自定义表名
func (Menu) TableName() string {
	return "sys_menu"
}

// Resource 资源表
type Resource struct {
	ID        string `gorm:"type:char(22);primaryKey" json:"id"`
	Name      string `gorm:"not null;size:64" json:"name"`
	Path      string `gorm:"size:128" json:"path"`
	Method    string `gorm:"size:16" json:"method"`
	ParentID  string `gorm:"type:char(22);" json:"parent_id"`
	CreatedAt int64  `gorm:"autoCreateTime:milli" json:"created_at"`
}

// TableName 自定义表名
func (Resource) TableName() string {
	return "sys_resource"
}
