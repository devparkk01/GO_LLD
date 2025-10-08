package main 

func main() {
	c := NewLoggerConfig(LogLevelError, NewConsoleAppender()) 
	logger := GetLogger(c)

	logger.Info("This is info message. it won't get logged.")
	logger.Error("This is error message. it will get logged.")
	logger.Fatal("This is fatal message. it will get logged.")

	logger.SetConfig(NewLoggerConfig(LogLevelDebug, NewFileAppender("log.txt")))
	logger.Info("This is info. it will get logged")
	logger.Error("Error message. it will get logged")
	logger.Debug("Debug message. it will get logged.")


}