package util

import "github.com/jinzhu/copier"

// CopyStruct 用于结构体间字段自动复制，常用于 req 到 dto、dto 到 model 等转换。
func CopyStruct(toValue interface{}, fromValue interface{}) error {
	return copier.Copy(toValue, fromValue)
}
