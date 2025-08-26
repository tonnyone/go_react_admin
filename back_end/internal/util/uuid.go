package util

import (
	"github.com/lithammer/shortuuid/v4"
)

// GenerateID 生成一个22位左右的唯一ID，无需init，适合用户ID、订单号等
func GenerateID() string {
	return shortuuid.New()
}
