package util

import (
	"path/filepath"
	"runtime"
)

/**
 * @description: 获取当前目录
 * @return {string}
 */
func CurDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filename)
}
