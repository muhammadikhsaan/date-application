package modules

import "sync"

type (
	cache struct {
		sync.Mutex
		data map[string]any
	}

	Cache interface {
		Set(key string, value any)
		Get(key string) any
		GetAll() map[string]any
		Delete(key string)
	}
)

func NewCache() Cache {
	return &cache{}
}

func (c *cache) Set(key string, value any) {
	c.Lock()
	defer c.Unlock()

	c.data[key] = value
}

func (c *cache) Get(key string) any {
	if value, ok := c.data[key]; ok {
		return value
	}

	return nil
}

func (c *cache) GetAll() map[string]any {
	return c.data
}

func (c *cache) Delete(key string) {
	c.Lock()
	defer c.Unlock()

	delete(c.data, key)
}
