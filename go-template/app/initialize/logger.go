package initialize

import (
	"fmt"
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
)

func initLogger() {
	encoder := getEncoder()

	// ErrorLevel以下归到info日志
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	// ErrorLevel及以上归到error日志
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	if exist, _ := utils.PathExists("logs/info/"); !exist {
		utils.CreateDir("logs/info/")
	}

	if exist, _ := utils.PathExists("logs/error/"); !exist {
		utils.CreateDir("logs/error/")
	}

	if exist, _ := utils.PathExists("logs/access/"); !exist {
		utils.CreateDir("logs/access/")
	}

	infoWriter := getLogWriter("logs/info/info")
	errorWriter := getLogWriter("logs/error/error")

	var coreTees []zapcore.Core
	if g.Cfg.Mode == "debug" {
		coreTees = append(coreTees, []zapcore.Core{
			zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoWriter, zapcore.AddSync(os.Stdout)), infoLevel),
			zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorWriter, zapcore.AddSync(os.Stdout)), errorLevel),
		}...)
	} else {
		coreTees = []zapcore.Core{
			zapcore.NewCore(encoder, infoWriter, infoLevel),
			zapcore.NewCore(encoder, errorWriter, errorLevel),
		}
	}

	core := zapcore.NewTee(coreTees...)

	//生成Logger
	g.Logger = zap.New(core, zap.AddCaller())
}

// getEncoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(g.TimeFormat)
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 得到LogWriter
func getLogWriter(filename string) zapcore.WriteSyncer {
	warnIoWriter := utils.GetRotateLogWriter(filename+"_%Y%m%d.log", fmt.Sprintf("logs/%s.log", path.Base(filename)))
	return zapcore.AddSync(warnIoWriter)
}
