package dto

// UserDTO 定义用户业务参数
type UserDTO struct {
	ID         string     `json:"id,omitempty"` // uuid主键
	Username   string     `json:"username,omitempty"`
	Email      string     `json:"email,omitempty"`
	Phone      string     `json:"phone,omitempty"`
	Password   string     `json:"password,omitempty"` // 忽略掉
	Department string     `json:"department,omitempty"`
	CreatedAt  int64      `json:"created_at,omitempty"`
	UpdatedAt  int64      `json:"updated_at,omitempty"`
	Disabled   bool       `json:"disabled,omitempty"`
	Roles      []UserRole `json:"roles,omitempty"` // 关联的角色
}

type UserRole struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
