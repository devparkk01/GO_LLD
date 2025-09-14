package main 

import (
	"fmt"
)

func main() {
	lru := &Lru{}
	cache := InitCache(lru, 4)
	cache.Add("something", "nothing")
	cache.Add("anyways", "noways")

	fifo := &Fifo{}

	cache.SetEvictionAlgo(fifo)
	cache.Add("hey", "bye")
	cache.Add("yoyo", "oye hoye ")

	value, ok := cache.Get("hey")
	if ok {
		fmt.Println("value " + value )
	}


}