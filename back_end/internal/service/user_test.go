package service

import (
	"context"
	"testing"

	"github.com/tonnyone/go_react_admin/internal/dao"
	"gorm.io/gorm"
)

func cleanTestUser(t *testing.T, db *gorm.DB, userDAO *dao.UserDAO, phone, email string) {
	err1 := userDAO.DeleteByPhone(context.Background(), db, phone)
	err2 := userDAO.DeleteByEmail(context.Background(), db, email)
	if err1 != nil || err2 != nil {
		t.Fatalf("failed to clean test user: %v %v", err1, err2)
	}
}

func TestUserService_RegisterByPhone(t *testing.T) {
	db := setupPostgresTestDB(t)
	userDAO := dao.NewUserDAO()
	userService := NewUserService(userDAO, db)
	phone := "18888888888"
	email := ""
	cleanTestUser(t, db, userDAO, phone, "")
	password := "pgtest123"
	err := userService.Register(context.Background(), phone, email, password)
	if err != nil {
		t.Fatalf("register by phone failed: %v", err)
	}
}

func TestUserService_RegisterByEmail(t *testing.T) {
	db := setupPostgresTestDB(t)
	userDAO := dao.NewUserDAO()
	userService := NewUserService(userDAO, db)
	phone := ""
	email := "pgtest@example.com"
	cleanTestUser(t, db, userDAO, "", email)
	password := "pgtest456"
	err := userService.Register(context.Background(), phone, email, password)
	if err != nil {
		t.Fatalf("register by email failed: %v", err)
	}
}

func TestUserService_LoginByPhone(t *testing.T) {
	db := setupPostgresTestDB(t)
	userDAO := dao.NewUserDAO()
	userService := NewUserService(userDAO, db)
	phone := "18888888888"
	password := "pgtest123"
	cleanTestUser(t, db, userDAO, phone, "")
	_ = userService.Register(context.Background(), phone, "", password)
	dto := &LoginDTO{Account: phone, Password: password}
	_, err := userService.Login(context.Background(), dto)
	if err != nil {
		t.Errorf("login by phone failed: %v", err)
	}
}

func TestUserService_LoginByEmail(t *testing.T) {
	db := setupPostgresTestDB(t)
	userDAO := dao.NewUserDAO()
	userService := NewUserService(userDAO, db)
	email := "pgtest@example.com"
	password := "pgtest456"
	cleanTestUser(t, db, userDAO, "", email)
	_ = userService.Register(context.Background(), "", email, password)
	dto := &LoginDTO{Account: email, Password: password}
	_, err := userService.Login(context.Background(), dto)
	if err != nil {
		t.Errorf("login by email failed: %v", err)
	}
}

func TestUserService_RegisterPhoneExists(t *testing.T) {
	db := setupPostgresTestDB(t)
	userDAO := dao.NewUserDAO()
	userService := NewUserService(userDAO, db)
	phone := "18888888888"
	cleanTestUser(t, db, userDAO, phone, "")
	_ = userService.Register(context.Background(), phone, "", "pgtest123")
	err := userService.Register(context.Background(), phone, "", "any")
	if err == nil || err.Error() != "手机号已存在" {
		t.Errorf("expected 手机号已存在 error, got: %v", err)
	}
}

func TestUserService_RegisterEmailExists(t *testing.T) {
	db := setupPostgresTestDB(t)
	userDAO := dao.NewUserDAO()
	userService := NewUserService(userDAO, db)
	email := "pgtest@example.com"
	cleanTestUser(t, db, userDAO, "", email)
	_ = userService.Register(context.Background(), "", email, "pgtest456")
	err := userService.Register(context.Background(), "", email, "any")
	if err == nil || err.Error() != "邮箱已存在" {
		t.Errorf("expected 邮箱已存在 error, got: %v", err)
	}
}
