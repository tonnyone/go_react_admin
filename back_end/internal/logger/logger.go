package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// Logger接口，便于后续切换日志库
type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	WithFields(fields map[string]interface{}) Entry
}

// Entry接口，兼容WithFields链式调用
type Entry interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
}

type logrusLogger struct {
	l *logrus.Logger
}

type logrusEntry struct {
	e *logrus.Entry
}

func (l *logrusLogger) Info(args ...interface{})                  { l.l.Info(args...) }
func (l *logrusLogger) Infof(format string, args ...interface{})  { l.l.Infof(format, args...) }
func (l *logrusLogger) Error(args ...interface{})                 { l.l.Error(args...) }
func (l *logrusLogger) Errorf(format string, args ...interface{}) { l.l.Errorf(format, args...) }
func (l *logrusLogger) Warn(args ...interface{})                  { l.l.Warn(args...) }
func (l *logrusLogger) Warnf(format string, args ...interface{})  { l.l.Warnf(format, args...) }
func (l *logrusLogger) Debug(args ...interface{})                 { l.l.Debug(args...) }
func (l *logrusLogger) Debugf(format string, args ...interface{}) { l.l.Debugf(format, args...) }
func (l *logrusLogger) WithFields(fields map[string]interface{}) Entry {
	return &logrusEntry{e: l.l.WithFields(fields)}
}

func (e *logrusEntry) Info(args ...interface{})                  { e.e.Info(args...) }
func (e *logrusEntry) Infof(format string, args ...interface{})  { e.e.Infof(format, args...) }
func (e *logrusEntry) Error(args ...interface{})                 { e.e.Error(args...) }
func (e *logrusEntry) Errorf(format string, args ...interface{}) { e.e.Errorf(format, args...) }
func (e *logrusEntry) Warn(args ...interface{})                  { e.e.Warn(args...) }
func (e *logrusEntry) Warnf(format string, args ...interface{})  { e.e.Warnf(format, args...) }
func (e *logrusEntry) Debug(args ...interface{})                 { e.e.Debug(args...) }
func (e *logrusEntry) Debugf(format string, args ...interface{}) { e.e.Debugf(format, args...) }

var stdLogger Logger = &logrusLogger{l: logrus.New()}

// Init 初始化日志配置，建议在main中调用
// level: info/debug/warn/error, format: text/json
func Init(level, format string) {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	switch format {
	case "json":
		l.SetFormatter(&logrus.JSONFormatter{})
	default:
		l.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	}
	switch level {
	case "debug":
		l.SetLevel(logrus.DebugLevel)
	case "warn":
		l.SetLevel(logrus.WarnLevel)
	case "error":
		l.SetLevel(logrus.ErrorLevel)
	default:
		l.SetLevel(logrus.InfoLevel)
	}
	stdLogger = &logrusLogger{l: l}
}

// 业务代码推荐只用以下接口
func Info(args ...interface{})                       { stdLogger.Info(args...) }
func Infof(format string, args ...interface{})       { stdLogger.Infof(format, args...) }
func Error(args ...interface{})                      { stdLogger.Error(args...) }
func Errorf(format string, args ...interface{})      { stdLogger.Errorf(format, args...) }
func Warn(args ...interface{})                       { stdLogger.Warn(args...) }
func Warnf(format string, args ...interface{})       { stdLogger.Warnf(format, args...) }
func Debug(args ...interface{})                      { stdLogger.Debug(args...) }
func Debugf(format string, args ...interface{})      { stdLogger.Debugf(format, args...) }
func WithFields(fields map[string]interface{}) Entry { return stdLogger.WithFields(fields) }

// UnderlyingLogger 仅供集成第三方中间件（如Gin日志）使用，业务代码请勿直接调用
func UnderlyingLogger() *logrus.Logger {
	if l, ok := stdLogger.(*logrusLogger); ok {
		return l.l
	}
	// 如果stdLogger不是*logrusLogger类型，返回一个默认的logrus实例
	// 或者根据你的错误处理策略，也可以panic
	return logrus.StandardLogger()
}

// Writer 仅供集成第三方中间件（如Gin日志）使用，业务代码请勿直接调用
func Writer() *io.PipeWriter {
	if l, ok := stdLogger.(*logrusLogger); ok {
		return l.l.Writer()
	}
	return logrus.StandardLogger().Writer()
}
