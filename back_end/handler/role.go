package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tonnyone/go_react_admin/internal/dao"
	"github.com/tonnyone/go_react_admin/internal/request"
	"github.com/tonnyone/go_react_admin/internal/service"
	"gorm.io/gorm"
)

// 新增角色
func CreateRoleHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req service.RoleReq
		if err := c.ShouldBindJSON(&req); err != nil {
			ResponseParamError(c, err)
			return
		}
		role := &dao.Role{
			Name:     req.Name,
			Describe: req.Describe,
		}
		service := service.NewRoleService()
		if err := service.CreateRole(c.Request.Context(), db, role); err != nil {
			ResponseFail(c, err.Error())
			return
		}
		ResponseSuccss(c, role)
	}
}

// 修改角色
func UpdateRoleHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var req service.RoleReq
		if err := c.ShouldBindJSON(&req); err != nil {
			ResponseParamError(c, err)
			return
		}
		service := service.NewRoleService()
		role := &dao.Role{
			Name:     req.Name,
			Describe: req.Describe,
			Type:     req.Type,
		}
		if err := service.UpdateRole(c.Request.Context(), db, id, role); err != nil {
			ResponseFail(c, err.Error())
			return
		}
		ResponseSuccss(c, role)
	}
}

// 删除角色
func DeleteRoleHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		service := service.NewRoleService()
		if err := service.DeleteRole(c.Request.Context(), db, id); err != nil {
			ResponseFail(c, err.Error())
			return
		}
		ResponseSuccss(c, true)
	}
}

// 分页查询角色
func ListRoleHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		pageReq := request.ParsePageReq(c)
		service := service.NewRoleService()
		roles, total, err := service.ListRoles(c.Request.Context(), db, pageReq.Page, pageReq.PageSize)
		if err != nil {
			ResponseFail(c, err.Error())
			return
		}
		ResponseSuccss(c, gin.H{
			"total":     total,
			"page":      pageReq.Page,
			"page_size": pageReq.PageSize,
			"roles":     roles,
		})
	}
}
