package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tonnyone/go_react_admin/internal/logger"
	"github.com/tonnyone/go_react_admin/internal/service"
	"gorm.io/gorm"
)

// BasicAuthMiddleware 基于数据库的 Basic Auth 认证中间件
func BasicAuthMiddleware(userService *service.UserService, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.Header("WWW-Authenticate", "Basic realm=Restricted")
			c.AbortWithStatus(401)
			return
		}
		logger.Info("basic auth pass")
		dto := &service.LoginDTO{Account: username, Password: password}
		_, err := userService.Login(c.Request.Context(), db, dto)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}
		c.Set("user", username)
		c.Next()
	}
}
