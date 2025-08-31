package dao

import (
	"context"

	"github.com/tonnyone/go_react_admin/internal/dto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	ID         string  `gorm:"type:char(36);primaryKey"` // uuid主键
	Username   string  `gorm:"not null;size:64"`
	Email      string  `gorm:"unique;size:128"`
	Phone      string  `gorm:"unique;size:64"`
	Password   string  `gorm:"not null;size:128" json:"password,omitempty"` // 忽略掉
	Department string  `gorm:"size:128"`
	CreatedAt  int64   `gorm:"autoCreateTime:milli"`
	UpdatedAt  int64   `gorm:"autoUpdateTime:milli"`
	Disabled   bool    `gorm:"default:false"`
	Roles      []*Role `gorm:"many2many:user_roles;"`
}

type UserRole struct {
	UserID    string `gorm:"type:char(36);not null;primaryKey"`
	RoleID    string `gorm:"type:char(36);not null;primaryKey"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	CreatedBy string `gorm:"type:char(36);not null;"`
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

// GetUsers 根据分页、排序和过滤条件获取用户列表
func (d *UserDAO) GetUsers(ctx context.Context, db *gorm.DB, pager *dto.Pager) ([]User, int64, error) {
	// 直接调用通用查询函数，并告诉它要排除 'Password' 字段
	return PaginatedQuery[User](ctx, db, pager, "Password")
}

// GetUser 根据ID获取用户信息
func (d *UserDAO) GetUser(ctx context.Context, db *gorm.DB, id string) (*User, error) {
	var user User
	err := db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// ClearUserRoles 清理用户角色
func ClearUserRoles(ctx context.Context, db *gorm.DB, userID string) error {
	return db.WithContext(ctx).Where("user_id = ?", userID).Delete(&UserRole{}).Error
}

// AppendUserRoles 为用户追加新的角色，如果角色已存在则忽略
func (d *UserDAO) AppendUserRoles(ctx context.Context, db *gorm.DB, userID string, roleIDs []string) error {
	if len(roleIDs) == 0 {
		return nil
	}
	var userRoles []UserRole
	for _, roleID := range roleIDs {
		userRoles = append(userRoles, UserRole{
			UserID: userID,
			RoleID: roleID,
		})
	}
	// 使用 OnConflict(DoNothing: true) 来避免因主键冲突而报错
	// 这会只插入那些用户尚不具备的角色关联
	return db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&userRoles).Error
}
