package cache

import (
	"time"
)

type Cache struct {
	storage map[string]ValueStorage
}

type ValueStorage struct {
	value    string
	deadline *time.Time
}

func NewCache() Cache {
	return Cache{
		storage: make(map[string]ValueStorage),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	value, ok := c.storage[key]
	if !ok {
		return "", false
	}
	if value.deadline != nil && !value.deadline.After(time.Now()) {
		delete(c.storage, key)
		return "", false
	}
	return value.value, ok
}

func (c *Cache) Put(key, value string) {
	c.storage[key] = ValueStorage{value: value, deadline: nil}
}

func (c *Cache) Keys() []string {
	var result []string
	currentTime := time.Now()

	for key, value := range c.storage {
		if value.deadline != nil && !value.deadline.After(currentTime) {
			delete(c.storage, key)
			continue
		}
		result = append(result, key)
	}
	return result
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.storage[key] = ValueStorage{value: value, deadline: &deadline}
}
