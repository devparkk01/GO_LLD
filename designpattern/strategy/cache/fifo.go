package main

import "fmt"

type Fifo struct {

}

func (f *Fifo) Evict(c *Cache) {
	fmt.Println("Evicting by fifo strategy.")	
}