package main

type (
	Cache[K comparable, V any] struct {
		data map[K]V
	}
)

// retorna a referência de um novo cache

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		data: make(map[K]V),
	}
}

// Pega o tipo da criação do cache e adiciona um novo item ao cache, usando o tipo de chave e valor genérico.

func (c *Cache[K, V]) Get(key K) (V, bool) {
	value, ok := c.data[key]
	return value, ok
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.data[key] = value
}

func main() {
	cache := NewCache[string, int]()

	cache.data["um"] = 1

	cache.Set("dois", 2)

	if value, ok := cache.Get("um"); ok {
		println("Valor de 'um':", value)
	}
	if value, ok := cache.Get("dois"); ok {
		println("Valor de 'dois':", value)
	}
}
