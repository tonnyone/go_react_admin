package service

import (
	"context"
	"testing"
	"time"

	"github.com/tonnyone/go_react_admin/internal/dao"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}
	if err := db.AutoMigrate(&dao.User{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return db
}

func TestUserService_Login(t *testing.T) {
	db := setupTestDB(t)
	userDAO := dao.NewUserDAO()
	userService := NewUserService(userDAO)

	// 插入测试用户
	db.Create(&dao.User{Username: "testuser", Password: "123456"})

	// 构造一个超时 context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()

	// 模拟慢查询（这里直接 sleep，实际可在 dao 层 sleep 或 mock）
	time.Sleep(2 * time.Millisecond)
	dto := &LoginDTO{Username: "testuser", Password: "123456"}
	token, err := userService.Login(ctx, db, dto)
	if err != nil || token == "" {
		t.Errorf("expected login success, got err: %v, token: %s", err, token)
	}

	// 错误密码
	dtoFail := &LoginDTO{Username: "testuser", Password: "wrong"}
	_, err = userService.Login(ctx, db, dtoFail)
	if err == nil {
		t.Errorf("expected login failure, got nil error")
	}
}
