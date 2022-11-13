package utils

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"log"
	"time"
)

// GetLogWriter 日志文件切割
func GetLogWriter(filename, linkname string) io.Writer {
	// 保存日志3天，每24小时分割一次日志
	hook, err := rotatelogs.New(
		filename,
		rotatelogs.WithLinkName(linkname),
		rotatelogs.WithMaxAge(time.Hour*24*3),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {
		log.Fatalf("rotatelogs error : %v", err)
	}
	return hook
}
