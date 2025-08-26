package database

import (
	"github.com/tonnyone/go_react_admin/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

// InitDB 初始化全局 DB 实例（可根据实际配置调整参数）
// logLevel: "info", "warn", "error", "silent", "debug"
func InitDB(dsn string, logLevel string) (*gorm.DB, error) {
	var err error
	dbOnce.Do(func() {
		gormLogger := logger.NewGormLogger(logLevel)
		dbInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: gormLogger,
		})
	})
	return dbInstance, err
}

// GetDB 获取全局 DB 实例
func GetDB() *gorm.DB {
	if dbInstance == nil {
		panic("DB not initialized, please call InitDB first")
	}
	return dbInstance
}
