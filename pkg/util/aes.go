package util

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

// MD5 md5函数
func MD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

// Base64Str base64字符串
func Base64Str(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
