package main

import "sync"

type Logger struct {
	config *LoggerConfig
	mu sync.RWMutex
}

// single instance of Logger. 
var (
	loggerInstance *Logger
	loggerOnce     sync.Once
)

func GetLogger(config *LoggerConfig) *Logger {
	loggerOnce.Do(func(){
		loggerInstance = &Logger{
			config: config,
		}
	})
	return loggerInstance
}


func(l *Logger) SetConfig(c *LoggerConfig) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.config = c 
}

func (l *Logger) Log(level LogLevel, message string) error {
	l.mu.RLock()
	configLevel := l.config.level
	appender := l.config.appender
	l.mu.RUnlock()
	// if config's level is greater than level of the message then it's ignored 
	if configLevel > level {
		return nil
	}

	logMessage := NewLogMessage(level, message)
	err := appender.append(logMessage)
	if err != nil {
		return err
	}
	return nil
}

func (l *Logger) Debug(message string) error {
	return l.Log(LogLevelDebug, message)
}

func (l *Logger) Info(message string) error {
	return l.Log(LogLevelInfo, message)
}
func (l *Logger) Warn(message string) error {
	return l.Log(LogLevelWarn, message)
}

func (l *Logger) Error(message string) error {
	return l.Log(LogLevelError, message)
}

func (l *Logger) Fatal(message string) error {
	return l.Log(LogLevelFatal, message)
}

