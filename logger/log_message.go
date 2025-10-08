package main

import (
	"fmt"
	"time"
)

type LogMessage struct {
	logLevel LogLevel
	message   string
	timestamp int64
}

func NewLogMessage(logLevel LogLevel, message string) *LogMessage {
	return &LogMessage{
		logLevel: logLevel,
		message: message,
		timestamp: time.Now().UnixMilli(),
	}
}

func(m *LogMessage) toString() string {
	return fmt.Sprintf("[%s] %d: %s", m.logLevel, m.timestamp, m.message)
}