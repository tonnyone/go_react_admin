package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tonnyone/go_react_admin/handler"
	"github.com/tonnyone/go_react_admin/internal/dao"
	"github.com/tonnyone/go_react_admin/internal/database"
	"github.com/tonnyone/go_react_admin/internal/service"
)

func RegisterRoutes(r *gin.Engine) {
	userDAO := dao.NewUserDAO()
	userService := service.NewUserService(userDAO)
	db := database.GetDB()

	r.POST("/login", handler.NewLoginHandler(userService, db))
	r.POST("/register", handler.NewRegisterHandler(userService, db))
	r.POST("/logout", handler.NewLogoutHandler())
}
