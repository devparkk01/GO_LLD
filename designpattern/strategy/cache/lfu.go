package main 

import "fmt"

type Lfu struct {

}

func (l *Lfu) Evict( c *Cache) {
	fmt.Println("Evicting by LFU strategy.")
	// for k, _ := range c.storage {
	// 	delete(c.storage, k)
	// 	break
	// }
}