package handler

import (
	"errors"
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
	Phone    string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}

func NewLoginHandler(userService *service.UserService, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			ResponseParamError(c, err)
			return
		}
		dto := service.LoginDTO{
			Account:  req.Username,
			Password: req.Password,
		}
		err := userService.Login(c.Request.Context(), db, &dto)
		if err != nil {
			ResponseFail(c, err.Error())
			return
		}
		ResponseSuccss(c, nil)
	}
}

// 注册
func NewRegisterHandler(userService *service.UserService, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			ResponseParamError(c, err)
			return
		}
		if req.Phone == "" && req.Email == "" {
			ResponseParamError(c, errors.New("手机号或邮箱必须填写一个"))
			return
		}
		if req.Password == "" {
			ResponseParamError(c, errors.New("密码不能为空"))
			return
		}
		err := userService.Register(c.Request.Context(), db, req.Phone, req.Email, req.Password)
		if err != nil {
			ResponseFail(c, err.Error())
			return
		}
		ResponseSuccss(c, nil)
	}
}

// 登出
func NewLogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: 清理token或session
		ResponseSuccss(c, nil)
	}
}
