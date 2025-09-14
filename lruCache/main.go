package main

import "fmt"

func main() {
	lruCache := NewLRUCache[int, string](3)

	value, ok := lruCache.Get(10)
	if ok {
		fmt.Println(value)
	} else {
		lruCache.Put(10, "TEN days")
		value, _ = lruCache.Get(10)
		fmt.Println(value)
	}
	lruCache.Put(20 , "Twenty bucks")
	lruCache.Put(20, "Twenty blokes")
	lruCache.DisplayCache()

	lruCache.Put(100, "Hundred tigers")
	lruCache.Get(10)
	lruCache.Put(1, "One above all")
	fmt.Println(lruCache.KeysMRUToLRU())


}