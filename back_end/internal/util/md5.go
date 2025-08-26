package util

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 返回字符串的md5值（不推荐生产环境，仅供演示/开发）
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
