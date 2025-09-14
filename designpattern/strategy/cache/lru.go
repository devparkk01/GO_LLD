package main

import "fmt"

type Lru struct {

}

func (l *Lru)Evict(c *Cache) {
	fmt.Println("Evicting by LRU strategy.")
	for k := range c.storage {
		delete(c.storage, k)
		break
	}
}
