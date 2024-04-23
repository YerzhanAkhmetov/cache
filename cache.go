package cache

import "fmt"

type Cache struct {
	data map[string]interface{}
}

func New() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

type CacheInterface interface {
	Set(string)
	Get(string) interface{}
	Delete(int)
	Show()
}

func (c *Cache) Set(key string, value int) {
	c.data[key] = value
}

func (c *Cache) Get(key string) interface{} {
	return c.data[key]
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}
func (c *Cache) Show() {
	fmt.Println(c.data)
}
