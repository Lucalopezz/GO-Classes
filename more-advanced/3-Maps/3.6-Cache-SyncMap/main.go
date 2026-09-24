package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	data sync.Map
}

// To not create a copy of the cache, we will return a pointer to it.

func NewCache() *Cache {
	return &Cache{}
}

// Reciver is a pointer to the cache, so we can modify the cache in place.

func (c *Cache) Set(key string, value any) {
	c.data.Store(key, value)
}

func (c *Cache) Get(key string) (any, bool) {
	result, ok := c.data.Load(key)
	return result, ok
}

func main() {
	var (
		cache = NewCache()
		wg    sync.WaitGroup
	)

	for i := 1; i <= 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.Set(fmt.Sprintf("key-%d", i), i)
		}()

	}
	for i := 1; i <= 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			key, ok := cache.Get(fmt.Sprintf("key-%d", i))
			if ok {
				fmt.Printf("key: %s, value: %v\n", fmt.Sprintf("key-%d", i), key)
			} else {
				fmt.Printf("key: %s, value: not found\n", fmt.Sprintf("key-%d", i))
			}
		}()
	}

	wg.Wait()
}
