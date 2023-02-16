package utils

import (
	"crypto/sha256"
	"fmt"
)

// 返回 SHA256 处理后的字符串
func GetSHA256Str(input string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(input)))
}
