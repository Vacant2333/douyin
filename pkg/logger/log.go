package logger

import (
	"sync"

	"go.uber.org/zap"
)

var zlog *zap.SugaredLogger
var once sync.Once

func init() {
	once.Do(func() {
		zlog = NewLogger()
		zlog = zlog.WithOptions(zap.AddCallerSkip(1))
	})
}

func NewLogger() *zap.SugaredLogger {
	// Todo: Production level
	level := zap.NewAtomicLevel()
	level.SetLevel(zap.DebugLevel)

	cfg := zap.NewProductionConfig()
	cfg.Level = level
	logger := zap.Must(cfg.Build())
	sugaredLogger := logger.Sugar()
	return sugaredLogger
}

func DebugF(format string, args ...interface{}) {
	zlog.Debugf(format, args...)
}

func InfoF(format string, args ...interface{}) {
	zlog.Infof(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	zlog.Fatalf(format, args...)
}

func Debug(args ...interface{}) {
	zlog.Debug(args...)
}

func Info(args ...interface{}) {
	zlog.Info(args...)
}

func Fatal(args ...interface{}) {
	zlog.Fatal(args...)
}
