package logger

import (
	"fmt"
)

func (l *Logger) sprintf(msg string, args ...interface{}) string {
	if len(args) == 0 {
		return msg
	}

	return fmt.Sprintf(msg, args...)
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.log.Debug(l.sprintf(msg))
}
func (l *Logger) Info(msg string, args ...interface{}) {
	l.log.Info(l.sprintf(msg, args...))
}
func (l *Logger) Warn(msg string, args ...interface{}) {
	l.log.Warn(l.sprintf(msg, args...))
}
func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.log.Fatal(l.sprintf(msg, args...))
}
func (l *Logger) Error(msg string, args ...interface{}) {
	l.log.Error(l.sprintf(msg, args...))
}
func (l *Logger) Panic(msg string, args ...interface{}) {
	l.log.Panic(l.sprintf(msg, args...))
}

// direct func
func Debug(msg string, args ...interface{}) {
	defaultLogger.Debug(msg, args...)
}
func Info(msg string, args ...interface{}) {
	defaultLogger.Info(msg, args...)
}
func Warn(msg string, args ...interface{}) {
	defaultLogger.Warn(msg, args...)
}
func Fatal(msg string, args ...interface{}) {
	defaultLogger.Fatal(msg, args...)
}
func Error(msg string, args ...interface{}) {
	defaultLogger.Error(msg, args...)
}
func Panic(msg string, args ...interface{}) {
	defaultLogger.Panic(msg, args...)
}
