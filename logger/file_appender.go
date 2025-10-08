package main

import (
	"os"
	"sync"
)

type FileAppender struct {
	filePath string
	mu       sync.Mutex
}

func NewFileAppender(filePath string) *FileAppender {
	return &FileAppender{
		filePath: filePath,
	}
}

func (f *FileAppender) append(message *LogMessage) error {
	
	f.mu.Lock()
	defer f.mu.Unlock()
	// O_APPEND : appending to the file, O_CREATE: create the file if it does not exist,
	// o_WRONLY: open the file in write mode only. 
	// 0644: Owner: 6(RWX), group: 4(RWX), others: 4(RWX)
	file, err := os.OpenFile(f.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err 
	}
	defer file.Close()

	_, err = file.WriteString(message.toString() + "\n")
	if err != nil {
		return err 
	}
	return nil 
}