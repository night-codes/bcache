package bcache

import (
	"sync"
	"time"
)

type (
	record struct {
		Time time.Time
		Data interface{}
	}

	// Cache is main cache struct
	Cache struct {
		sync.Mutex
		list    map[string]record
		updater func(key string) interface{}
	}
)

// Create new cache
func Create() Cache {
	c := Cache{
		list:    map[string]record{},
		updater: func(key string) interface{} { return nil },
	}
	go func() {
		for range time.Tick(100 * time.Millisecond) {
			for i, el := range c.list {
				if el.Time.Before(time.Now()) {
					delete(c.list, i)
				}
			}
		}
	}()
	return c
}

func (c *Cache) set(key string, data interface{}) {
	t := time.Now().Add(time.Second)
	if c.list[key].Data != nil {
		t = c.list[key].Time.Add(time.Second)
	}
	c.list[key] = record{
		Time: t,
		Data: data,
	}
}

// Set data to the bcache
func (c *Cache) Set(key string, data interface{}) {
	c.Lock()
	defer c.Unlock()
	c.set(key, data)
}

func (c *Cache) get(key string) interface{} {
	if c.list[key].Data == nil {
		data := c.updater(key)
		if data == nil {
			return nil
		}
		c.set(key, data)
		return data
	}

	c.list[key] = record{
		Time: c.list[key].Time.Add(time.Second),
		Data: c.list[key].Data,
	}
	return c.list[key].Data
}

// Get data from the bcache
func (c *Cache) Get(key string) interface{} {
	c.Lock()
	defer c.Unlock()
	return c.get(key)
}

// Updater is set function to update cache by key
func (c *Cache) Updater(fn func(key string) interface{}) {
	c.Lock()
	defer c.Unlock()
	c.updater = fn
}

// Clear cache data
func (c *Cache) Clear() {
	c.Lock()
	defer c.Unlock()
	c.list = map[string]record{}
}
