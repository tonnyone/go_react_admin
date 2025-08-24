package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tonnyone/go_react_admin/internal/service"
	"gorm.io/gorm"
)

// 登录请求结构体
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 注册请求结构体
type RegisterRequest struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Email      string `json:"email"`
	Department string `json:"department"`
}

func NewLoginHandler(userService *service.UserService, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			ParamError(c, err)
			return
		}
		dto := service.LoginDTO{
			Username: req.Username,
			Password: req.Password,
		}
		token, err := userService.Login(c.Request.Context(), db, &dto)
		if err != nil {
			Fail(c, err.Error())
			return
		}
		Success(c, gin.H{"token": token})
	}
}

// 注册
func NewRegisterHandler(userService *service.UserService, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			ParamError(c, err)
			return
		}
		// TODO: 检查用户名是否存在，写入数据库
		Success(c, "注册成功")
	}
}

// 登出
func NewLogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: 清理token或session
		Success(c, "已登出")
	}
}
