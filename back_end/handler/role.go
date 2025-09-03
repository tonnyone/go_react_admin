package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tonnyone/go_react_admin/internal/dao"
	"github.com/tonnyone/go_react_admin/internal/dto"
	"github.com/tonnyone/go_react_admin/internal/service"
)

// 新增角色
func CreateRoleHandler(roleService service.RoleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.RoleDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			ResponseParamError(c, err)
			return
		}
		if err := roleService.CreateRole(c.Request.Context(), &req); err != nil {
			ResponseFail(c, err.Error())
			return
		}
		ResponseSuccss[any](c, nil)
	}
}

// 修改角色
func UpdateRoleHandler(roleService service.RoleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var req dto.RoleDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			ResponseParamError(c, err)
			return
		}
		role := &dao.Role{
			Name:     req.Name,
			Describe: req.Describe,
			Type:     string(req.Type),
		}
		if err := roleService.UpdateRole(c.Request.Context(), id, role); err != nil {
			ResponseFail(c, err.Error())
			return
		}
		ResponseSuccss(c, role)
	}
}

// 删除角色
func DeleteRoleHandler(roleService service.RoleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := roleService.DeleteRole(c.Request.Context(), id); err != nil {
			ResponseFail(c, err.Error())
			return
		}
		ResponseSuccss(c, true)
	}
}

// 分页查询角色
func ListRoleHandler(roleService service.RoleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		pager, err := ParsePageReq(c)
		if err != nil {
			ResponseParamError(c, err)
			return
		}
		roles, total, err := roleService.ListRoles(c.Request.Context(), pager.Current, pager.PageSize)
		if err != nil {
			ResponseFail(c, err.Error())
			return
		}
		ResponseSuccss(c, PageData[dao.Role]{
			List:  roles,
			Total: total,
		})
	}
}
