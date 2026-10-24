package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	data sync.Map
}

// Para não criar uma cópia do cache, retornamos um ponteiro para ele.

func NewCache() *Cache {
	return &Cache{}
}

// O receptor é um ponteiro para o cache, então podemos modificá-lo diretamente.

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
				fmt.Printf("chave: %s, valor: %v\n", fmt.Sprintf("key-%d", i), key)
			} else {
				fmt.Printf("chave: %s, valor: não encontrado\n", fmt.Sprintf("key-%d", i))
			}
		}()
	}

	wg.Wait()
}
