package main 

type Cache struct {
	storage map[string]string
	evictStrategy EvictionAlgo
	maxCapacity int 
	curCapacity int 
}

func InitCache(evictStrategy EvictionAlgo, maxCapacity int) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage: storage,
		evictStrategy: evictStrategy,
		maxCapacity: maxCapacity,
		curCapacity: 0,
	}
}

func (c *Cache) SetEvictionAlgo(evictStrategy EvictionAlgo) {
	c.evictStrategy = evictStrategy
}

func (c *Cache) Add(key, value string) {
	if c.curCapacity >= c.maxCapacity {
		c.evictStrategy.Evict(c)
	}
	c.curCapacity++
	c.storage[key] = value 
}

func(c *Cache) Get(key string) (string, bool) {
	val, ok := c.storage[key]
	return val, ok
}