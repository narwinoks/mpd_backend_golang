package cache

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type Item struct {
	Value      []byte
	Expiration int64
}

type CacheStore struct {
	items map[string]Item
	mu    sync.RWMutex
	rdb   *redis.Client
}

var (
	instance *CacheStore
	once     sync.Once
	ctx      = context.Background()
)

func GetCache() *CacheStore {
	once.Do(func() {
		instance = &CacheStore{
			items: make(map[string]Item),
		}
	})
	return instance
}

func (c *CacheStore) SetRedisClient(rdb *redis.Client) {
	c.rdb = rdb
}

func (c *CacheStore) Set(key string, value interface{}, duration time.Duration) {
	data, err := json.Marshal(value)
	if err != nil {
		return
	}

	if c.rdb != nil {
		c.rdb.Set(ctx, key, data, duration)
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = Item{
		Value:      data,
		Expiration: time.Now().Add(duration).UnixNano(),
	}
}

func (c *CacheStore) Get(key string, target interface{}) bool {
	if c.rdb != nil {
		val, err := c.rdb.Get(ctx, key).Result()
		if err != nil {
			return false
		}
		err = json.Unmarshal([]byte(val), target)
		return err == nil
	}

	c.mu.RLock()
	item, found := c.items[key]
	c.mu.RUnlock()

	if !found {
		return false
	}
	if time.Now().UnixNano() > item.Expiration {
		c.Delete(key)
		return false
	}

	err := json.Unmarshal(item.Value, target)
	return err == nil
}

func (c *CacheStore) Delete(key string) {
	if c.rdb != nil {
		c.rdb.Del(ctx, key)
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}
