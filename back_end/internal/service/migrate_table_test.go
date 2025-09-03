package service

import (
	"testing"

	"github.com/tonnyone/go_react_admin/internal/config"
	"github.com/tonnyone/go_react_admin/internal/dao"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupPostgresTestDB(t *testing.T) *gorm.DB {
	cfg, err := config.LoadConfigWithPath("../../")
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}
	dsn := cfg.DB.DSN
	if dsn == "" {
		t.Skip("config DB.DSN not set, skip Postgres test")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open postgres db: %v", err)
	}
	return db
}

func TestRoleService_EnsureRoleTable(t *testing.T) {
	db := setupPostgresTestDB(t)
	err := db.AutoMigrate(
		&dao.User{},
		&dao.Role{},
		&dao.UserRole{},
	)
	if err != nil {
		t.Fatal("failed to ensure role table: %w", err)
	}
}
