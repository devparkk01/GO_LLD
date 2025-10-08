package main 

type LogLevel int 

const (
	LogLevelTrace LogLevel = iota
	LogLevelDebug
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

func (l LogLevel) String() string {
	return [...]string{"TRACE", "DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}[l]
}