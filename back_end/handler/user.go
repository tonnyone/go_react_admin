package handler

import (
	"errors"
	"github.com/gin-gonic/gin"

	"github.com/tonnyone/go_react_admin/internal/dto"
	"github.com/tonnyone/go_react_admin/internal/service"
)

// UserHandler 负责处理用户相关的HTTP请求
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler 创建一个新的 UserHandler 实例
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

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

func NewLoginHandler(userService *service.UserService) gin.HandlerFunc {
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
		token, err := userService.Login(c.Request.Context(), &dto)
		if err != nil {
			ResponseFail(c, err.Error())
			return
		}
		ResponseSuccss(c, token)
	}
}

// 注册
func NewRegisterHandler(userService *service.UserService) gin.HandlerFunc {
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
		err := userService.Register(c.Request.Context(), req.Phone, req.Email, req.Password)
		if err != nil {
			ResponseFail(c, err.Error())
			return
		}
		ResponseSuccssNoData(c)
	}
}

// NewGetUsersHandler 创建获取用户列表的处理器
func NewGetUsersHandler(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		pager, err := ParsePageReq(c)
		if err != nil {
			ResponseParamError(c, err)
			return
		}
		users, total, err := userService.GetUsers(c.Request.Context(), &pager)
		if err != nil {
			ResponseFail(c, "获取用户列表失败: "+err.Error())
			return
		}
		ResponseSuccss(c, PageData[dto.UserDTO]{
			List:  users,
			Total: total,
		})
	}
}

// 登出
func NewLogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: 清理token或session
		ResponseSuccssNoData(c)
	}
}

// BindRoles 处理为用户绑定角色的HTTP请求
func NewBindRolesHandler(service *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.BindUserRolesReq
		if err := c.ShouldBindJSON(&req); err != nil {
			ResponseParamError(c, err)
			return
		}
		if err := service.BindRolesToUser(c.Request.Context(), &req); err != nil {
			ResponseFail(c, "用户绑定角色失败: "+err.Error())
			return
		}
		ResponseSuccssNoData(c)
	}
}
