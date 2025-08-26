package dao

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	ID         string `gorm:"type:char(36);primaryKey"` // uuid主键
	Username   string `gorm:"not null;size:64"`
	Email      string `gorm:"unique;size:128"`
	Phone      string `gorm:"unique;size:64"`
	Password   string `gorm:"not null;size:128"`
	Department string `gorm:"size:128"`
	CreatedAt  int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt  int64  `gorm:"autoUpdateTime:milli"`
	Disabled   bool   `gorm:"default:false"`
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

// CheckUserByPhoneOrEmailExist 检查手机号或邮箱和密码是否存在
func (dao *UserDAO) CheckUserByPhoneOrEmailExist(ctx context.Context, db *gorm.DB, account, password string) (bool, error) {
	var user User
	err := db.WithContext(ctx).Where(
		"(phone = ? OR email = ?) AND password = ?",
		account, account, password,
	).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// CheckPhoneExist 检查手机号是否已注册
func (dao *UserDAO) CheckPhoneExist(ctx context.Context, db *gorm.DB, phone string) (bool, error) {
	var user User
	err := db.WithContext(ctx).Where("phone = ?", phone).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// CheckEmailExist 检查邮箱是否已注册
func (dao *UserDAO) CheckEmailExist(ctx context.Context, db *gorm.DB, email string) (bool, error) {
	var user User
	err := db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// DeleteByPhone 删除指定手机号的用户
func (dao *UserDAO) DeleteByPhone(ctx context.Context, db *gorm.DB, phone string) error {
	return db.WithContext(ctx).Where("phone = ?", phone).Delete(&User{}).Error
}

// DeleteByEmail 删除指定邮箱的用户
func (dao *UserDAO) DeleteByEmail(ctx context.Context, db *gorm.DB, email string) error {
	return db.WithContext(ctx).Where("email = ?", email).Delete(&User{}).Error
}
