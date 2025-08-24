package dao

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Username   string `gorm:"unique;not null;size:64"`
	Password   string `gorm:"not null;size:128"`
	Email      string `gorm:"size:128"`
	Department string `gorm:"size:128"`
	Disabled   bool   `gorm:"default:false"`
	Deleted    bool   `gorm:"default:false"`
	CreatedAt  int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt  int64  `gorm:"autoUpdateTime:milli"`
}

type UserDAO struct{}

// NewUserDAO 创建 UserDAO 实例
func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

// CheckUserExist 检查用户名和密码是否存在，db为必选参数（支持事务/全局DB）
func (dao *UserDAO) CheckUserExist(ctx context.Context, db *gorm.DB, username, password string) (bool, error) {
	var user User
	err := db.WithContext(ctx).Where("username = ? AND password = ?", username, password).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
