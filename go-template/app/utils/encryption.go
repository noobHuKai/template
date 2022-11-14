package utils

import (
	"crypto/md5"
	"fmt"
)

func MD5Encrypt(data string) string {
	has := md5.Sum([]byte(data))
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
