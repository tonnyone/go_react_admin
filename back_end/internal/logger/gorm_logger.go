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
	return &GormLogger{
		level: mapLogLevel(level),
	}
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
	// To see all SQL, the GORM log level must be `Info`.
	// We check this level before proceeding.
	if l.level < gormlogger.Info {
		return
	}

	elapsed := time.Since(begin)
	msg, rows := fc()

	// If there is an error, we log it as an Error, regardless of level (as long as it's not Silent).
	if err != nil && l.level >= gormlogger.Error {
		// Using a fields-based approach for structured logging
		WithFields(map[string]interface{}{
			"module":   "gorm",
			"duration": elapsed,
			"rows":     rows,
			"error":    err,
		}).Errorf(msg)
		return
	}

	// For slow queries, we log it as a Warning.
	// (You can configure the slow query threshold in gorm.Config)
	// For now, we'll just log all queries at Info level.

	// Log successful queries at Info level.
	if l.level >= gormlogger.Info {
		WithFields(map[string]interface{}{
			"module":   "gorm",
			"duration": elapsed,
			"rows":     rows,
		}).Infof(msg)
	}
}
