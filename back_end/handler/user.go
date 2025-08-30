package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/tonnyone/go_react_admin/internal/dao"
	"github.com/tonnyone/go_react_admin/internal/service"
	"gorm.io/gorm"
)

// 登录请求结构体
type LoginRequest struct {
	Account  string `json:"account" binding:"required"`
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
			Account:  req.Account,
			Password: req.Password,
		}
		token, err := userService.Login(c.Request.Context(), db, &dto)
		if err != nil {
			ResponseFail(c, err.Error())
			return
		}
		ResponseSuccss(c, token)
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
		ResponseSuccss[any](c, nil)
	}
}

// NewGetUsersHandler 创建获取用户列表的处理器
func NewGetUsersHandler(userService *service.UserService, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		pager, err := ParsePageReq(c)
		if err != nil {
			ResponseParamError(c, err)
			return
		}
		users, total, err := userService.GetUsers(c.Request.Context(), db, &pager)
		if err != nil {
			ResponseFail(c, "获取用户列表失败: "+err.Error())
			return
		}
		ResponseSuccss(c, PageData[dao.User]{
			List:  users,
			Total: total,
		})
	}
}

// 登出
func NewLogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: 清理token或session
		ResponseSuccss[any](c, nil)
	}
}
