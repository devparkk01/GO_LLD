package main 

type EvictionAlgo interface {
	Evict(c *Cache)
}

