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
	db := database.GetDB()
	userService := service.NewUserService(dao.NewUserDAO(), db)
	roleService := service.NewRoleService(dao.NewRoleDAO(), db)

	// --- 公共路由组 (无需认证) ---
	publicGroup := r.Group("")
	{
		publicGroup.POST("/login", handler.NewLoginHandler(userService))
		publicGroup.POST("/register", handler.NewRegisterHandler(userService))
	}

	// --- API 路由组 (需要认证) ---
	apiGroup := r.Group("")
	apiGroup.Use(middleware.BasicAuthMiddleware(userService))
	{
		// 用户管理
		userGroup := apiGroup.Group("user")
		{
			userGroup.GET("", handler.NewGetUsersHandler(userService))
			userGroup.PUT("/:id/bind_role", handler.NewBindRolesHandler(userService))
			// 在这里可以继续添加其他用户相关的路由，如：
			// userGroup.GET("/:id", handler.GetUserHandler(userService))
			// userGroup.POST("", handler.CreateUserHandler(userService))
		}

		// 角色管理
		roleGroup := apiGroup.Group("/role")
		{
			roleGroup.POST("", handler.CreateRoleHandler(*roleService)) // 假设这是获取角色列表
			roleGroup.GET("", handler.ListRoleHandler(*roleService))    // 假设这是获取角色列表
			// 在这里可以继续添加其他角色相关的路由
		}
	}
}
