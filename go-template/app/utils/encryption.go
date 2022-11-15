package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Encrypt(data []byte, b ...byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(b))
}
