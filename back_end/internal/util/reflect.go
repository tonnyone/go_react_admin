// internal/util/reflect.go
package util

import (
	"reflect"
	"strings"
	"sync"

	"github.com/iancoleman/strcase"
)

var fieldCache = sync.Map{}

// BuildAllowedFieldsFor a泛型函数，为任何结构体类型构建字段白名单
// 它直接接收类型参数 T，无需传递实例
func BuildAllowedFieldsFor[T any]() map[string]bool {
	// 1. 获取泛型参数 T 的类型
	var modelInstance T
	t := reflect.TypeOf(modelInstance)

	// 2. 尝试从缓存中获取
	if cachedFields, ok := fieldCache.Load(t); ok {
		return cachedFields.(map[string]bool)
	}

	// 3. 如果缓存未命中，则通过反射构建 (逻辑与之前完全相同)
	allowedFields := make(map[string]bool)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		gormTag := field.Tag.Get("gorm")
		if gormTag != "" && gormTag != "-" {
			if strings.HasPrefix(gormTag, "column:") {
				columnName := strings.Split(strings.Split(gormTag, ";")[0], ":")[1]
				allowedFields[columnName] = true
				continue
			}
		}
		fieldNameSnake := strcase.ToSnake(field.Name)
		allowedFields[fieldNameSnake] = true
	}

	// 4. 将结果存入缓存
	fieldCache.Store(t, allowedFields)

	return allowedFields
}
