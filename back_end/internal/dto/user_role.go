package dto

// BindUserRolesReq 定义了用户绑定角色请求的数据结构
type BindUserRolesReq struct {
	UserID  string   `json:"user_id" binding:"required"`
	RoleIDs []string `json:"role_ids" binding:"required"`
}
