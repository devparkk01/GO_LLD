package main

import (
	"fmt"
	"sync"
)

type ConsoleAppender struct {
	mu sync.Mutex
}

func NewConsoleAppender() *ConsoleAppender{
	return &ConsoleAppender{}
}

func (c *ConsoleAppender) append(message *LogMessage) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println(message.toString())
	return nil 
}