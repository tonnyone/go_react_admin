package logger

import (
	"context"
	"time"

	gormlogger "gorm.io/gorm/logger"
)

// GormLogger 适配 gorm.io/gorm/logger.Interface，内部调用 stdLogger
// 保证所有日志统一走 logger.go

type GormLogger struct {
	level gormlogger.LogLevel
}

func NewGormLogger(level string) *GormLogger {
	return &GormLogger{level: mapLogLevel(level)}
}

// mapLogLevel 将配置的日志级别映射为gormlogger.LogLevel
func mapLogLevel(level string) gormlogger.LogLevel {
	switch level {
	case "silent":
		return gormlogger.Silent
	case "error":
		return gormlogger.Error
	case "warn":
		return gormlogger.Warn
	case "info", "debug":
		return gormlogger.Info
	default:
		return gormlogger.Warn
	}
}

func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	l.level = level
	return l
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= gormlogger.Info {
		Infof("[GORM] "+msg, data...)
	}
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= gormlogger.Warn {
		Warnf("[GORM] "+msg, data...)
	}
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= gormlogger.Error {
		Errorf("[GORM] "+msg, data...)
	}
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.level == gormlogger.Silent {
		return
	}
	elapsed := time.Since(begin)
	msg, rows := fc()
	if err != nil && l.level >= gormlogger.Error {
		Errorf("[GORM] %s | %v | %d rows | %v", msg, elapsed, rows, err)
	} else if l.level >= gormlogger.Info {
		Debugf("[GORM] %s | %v | %d rows", msg, elapsed, rows)
	}
}
