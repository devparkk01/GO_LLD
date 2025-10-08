package main 

type LogAppender interface {
	append(message *LogMessage) error
}