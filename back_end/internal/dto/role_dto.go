package dto

// RoleDTO 用于角色相关的请求参数 DTO
// 只做参数传递，不含业务逻辑

type RoleDTO struct {
	Name     string   `json:"name" binding:"required"`
	Describe string   `json:"describe"`
	Type     RoleType `json:"type"`
}

type RoleType string

const (
	RoleTypeSystem RoleType = "system"
	RoleTypeCustom RoleType = "custom"
)
