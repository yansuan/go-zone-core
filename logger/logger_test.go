package logger

import "testing"

func TestMain(t *testing.T) {
	Debug("debug msg")
	Info("info msg")
	Warn("warn msg")
	Error("error msg")
	Fatal("fatal msg")
	Panic("panic msg")
}
