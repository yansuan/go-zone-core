package logger

var defaultLogger *Logger

func init() {
	Init()
}

func Init(options ...Option) {
	var option Option
	if len(options) == 0 {
		option.InfoFile = "log/info.log"
		option.ErrorFile = "log/error.log"
		option.MaxSize = 10
		option.MaxBackups = 7
		option.MaxAge = 30
		option.Compress = false
	}
	defaultLogger = New(option)
}
