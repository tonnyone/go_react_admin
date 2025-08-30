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
	publicGroup := r.Group("")
	{
		publicGroup.POST("/login", handler.NewLoginHandler(userService, db))
		publicGroup.POST("/register", handler.NewRegisterHandler(userService, db))
	}

	// --- API 路由组 (需要认证) ---
	apiGroup := r.Group("")
	apiGroup.Use(middleware.BasicAuthMiddleware(userService, db))
	{
		// 用户管理
		userGroup := apiGroup.Group("user")
		{
			userGroup.GET("/list", handler.NewGetUsersHandler(userService, db))
			// 在这里可以继续添加其他用户相关的路由，如：
			// userGroup.GET("/:id", handler.GetUserHandler(userService))
			// userGroup.POST("", handler.CreateUserHandler(userService))
		}

		// 角色管理
		roleGroup := apiGroup.Group("/role")
		{
			roleGroup.POST("/list", handler.ListRoleHandler(db)) // 假设这是获取角色列表
			// 在这里可以继续添加其他角色相关的路由
		}
	}
}
