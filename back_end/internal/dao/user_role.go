package dao

import (
	"context"
	"gorm.io/gorm"
)

// UserRole 关联表
type UserRole struct {
	ID        string `gorm:"type:char(36);primaryKey"`
	UserID    string `gorm:"type:char(36);not null;index"`
	RoleID    string `gorm:"type:char(36);not null;index"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
}

// UserRoleDAO 关联表操作
type UserRoleDAO struct{}

func NewUserRoleDAO() *UserRoleDAO {
	return &UserRoleDAO{}
}

// AddUserRole 新增用户-角色关联
func (dao *UserRoleDAO) AddUserRole(ctx context.Context, db *gorm.DB, userID, roleID, id string) error {
	ur := UserRole{ID: id, UserID: userID, RoleID: roleID}
	return db.WithContext(ctx).Create(&ur).Error
}

// DeleteUserRole 删除用户-角色关联
func (dao *UserRoleDAO) DeleteUserRole(ctx context.Context, db *gorm.DB, userID, roleID string) error {
	return db.WithContext(ctx).Where("user_id = ? AND role_id = ?", userID, roleID).Delete(&UserRole{}).Error
}

// ListRolesByUser 查询用户所有角色ID
func (dao *UserRoleDAO) ListRolesByUser(ctx context.Context, db *gorm.DB, userID string) ([]string, error) {
	var urs []UserRole
	if err := db.WithContext(ctx).Where("user_id = ?", userID).Find(&urs).Error; err != nil {
		return nil, err
	}
	ids := make([]string, 0, len(urs))
	for _, ur := range urs {
		ids = append(ids, ur.RoleID)
	}
	return ids, nil
}
