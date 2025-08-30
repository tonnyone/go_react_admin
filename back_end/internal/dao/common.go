package dao

import (
	"context"
	"fmt"

	"github.com/tonnyone/go_react_admin/internal/dto"
	"github.com/tonnyone/go_react_admin/internal/logger"
	"github.com/tonnyone/go_react_admin/internal/util"
	"gorm.io/gorm"
)

// T: GORM 模型类型
// PaginatedQuery 是一个通用的分页查询函数
// pager: 包含分页、过滤、排序信息的请求对象
// omitFields: 需要从查询结果中排除的字段列表
func PaginatedQuery[T any](ctx context.Context, db *gorm.DB, pager *dto.Pager, omitFields ...string) ([]T, int64, error) {
	var results []T
	var total int64

	tx := db.WithContext(ctx).Model(new(T))
	// 2. 获取模型的安全字段白名单，防止SQL注入
	allowedFields := util.BuildAllowedFieldsFor[T]()

	// 3. 应用过滤条件
	if pager.Filter != nil {
		for key, value := range pager.Filter {
			if allowedFields[key] {
				// 确保值不为空字符串
				if key != "" && value != "" {
					tx = tx.Where(fmt.Sprintf("%s LIKE ?", key), "%"+value+"%")
				}
			} else {
				logger.Warnf("通用查询忽略了无效的过滤字段: %s", key)
			}
		}
	}

	// 4. 获取总记录数 (在应用分页和排序之前)
	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 5. 应用排序条件
	if pager.Sorter != nil {
		for key, value := range pager.Sorter {
			if allowedFields[key] {
				order := "ASC"
				if value == dto.Desc {
					order = "DESC"
				}
				tx = tx.Order(fmt.Sprintf("%s %s", key, order))
			} else {
				logger.Warnf("通用查询忽略了无效的排序字段: %s", key)
			}
		}
	} else {
		// 提供一个默认排序，防止数据库因无序而导致分页结果不一致
		tx = tx.Order("created_at DESC")
	}

	// 6. 应用 Omit 和分页
	page := 1
	if pager.Current > 0 {
		page = pager.Current
	}

	pageSize := 10
	if pager.PageSize > 0 {
		if pager.PageSize > 100 {
			pageSize = 100
		} else {
			pageSize = pager.PageSize
		}
	}

	offset := (page - 1) * pageSize
	err := tx.Offset(offset).Limit(pageSize).
		Omit(omitFields...).
		Find(&results).Error
	if err != nil {
		return nil, 0, err
	}
	return results, total, nil
}
