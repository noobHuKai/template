package utils

import (
	"os"
)

func DeLFile(filePath string) error {
	return os.RemoveAll(filePath)
}

// FileExist 判断文件是否存在
func FileExist(path string) bool {
	fi, err := os.Lstat(path)
	if err == nil {
		return !fi.IsDir()
	}
	return !os.IsNotExist(err)
}
