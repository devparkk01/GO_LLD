package main 

type LoggerConfig struct {
	level LogLevel
	appender LogAppender
}

func NewLoggerConfig(level LogLevel, appender LogAppender) *LoggerConfig{
	return &LoggerConfig{
		level: level, 
		appender: appender,
	}
}