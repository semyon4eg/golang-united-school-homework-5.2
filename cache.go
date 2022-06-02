package cache

import (
	"time"
)

type Cache struct {
	storage map[string]string
	timer map[string]time.Time
}

func NewCache() Cache {
	return Cache{
		storage: make(map[string]string),
		timer: make(map[string]time.Time),
	}
}

func (cache Cache) Get(key string) (string, bool) {

	cache.CheckTimer()

	value, ok := cache.storage[key]
	return value, ok
}

func (cache Cache) Put(key, value string) {
	cache.storage[key] = value
}

func (cache Cache) Keys() []string {

	cache.CheckTimer()

	keys := []string{}
    for key := range cache.storage {
        keys = append(keys, key)
    }

	return keys
}

func (cache Cache) PutTill(key, value string, deadline time.Time) {
	cache.storage[key] = value
	cache.timer[key] = deadline
}

func (cache Cache) CheckTimer() {
	for key := range cache.storage {
		if deadline, ok := cache.timer[key]; ok {

			if deadline.Before(time.Now()) {
				delete(cache.storage, key)
				delete(cache.timer, key)
			}
		}
	}
}
