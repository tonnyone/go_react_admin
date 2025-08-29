package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tonnyone/go_react_admin/handler"
	"github.com/tonnyone/go_react_admin/internal/dao"
	"github.com/tonnyone/go_react_admin/internal/database"
	"github.com/tonnyone/go_react_admin/internal/middleware"
	"github.com/tonnyone/go_react_admin/internal/service"
)

func RegisterRoutes(r *gin.Engine) {
	// --- 依赖注入 ---
	userService := service.NewUserService(dao.NewUserDAO())
	db := database.GetDB()
	// --- 公共路由组 (无需认证) ---
	r.POST("/login", handler.NewLoginHandler(userService, db))
	r.POST("/register", handler.NewRegisterHandler(userService, db))

	// --- API 路由组 (需要认证) ---
	apiGroup := r.Group("/")
	// 在这里可以为整个 apiGroup 添加认证中间件
	apiGroup.Use(middleware.BasicAuthMiddleware(userService, db))
	// 用户管理
	userGroup := apiGroup.Group("user")
	// 假设 handler 中有对应的处理函数，如果没有请先创建
	userGroup.POST("list", handler.ListRoleHandler(db))
	// 角色管理
	roleGroup := apiGroup.Group("/role")
	roleGroup.POST("list2", handler.ListRoleHandler(db))
}
