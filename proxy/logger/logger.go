package logger

import (
	"time"

	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	logPath   = "D:/mygo/src/imooc/分布式爬虫项目/demo1/proxy/logs/demo1.log"
	logSize   = 100
	logBackUp = 2
	logMaxAge = 7
)

var Logger *zap.SugaredLogger

func init() {
	InitLogger(logPath, logSize, logBackUp, logMaxAge)
}

func InitLogger(logPath string, logSize int, logBackUp int, logMaxAge int) {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    logSize, // megabytes
		MaxBackups: logBackUp,
		MaxAge:     logMaxAge, // days
	})
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(NewEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),
			w),
		zap.DebugLevel,
	)
	Logger = zap.New(core, zap.AddCaller()).Sugar()

	Logger.Debug("init logger  success.")
}

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "log",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "trace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// func Debugf(template string, args ...interface{}) {
// 	Logger.Debugf(template, args)
// }

// func Infof(template string, args ...interface{}) {
// 	Logger.Infof(template, args)
// }

// func Warnf(template string, args ...interface{}) {
// 	Logger.Warnf(template, args)
// }

// func Errorf(template string, args ...interface{}) {
// 	Logger.Errorf(template, args)
// }
