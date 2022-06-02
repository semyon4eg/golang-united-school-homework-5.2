package cache

import (
	"time"
	"fmt"
)

type Cache struct {
	storage, timer map[string]string
}

func NewCache() Cache {
	return Cache{
		storage: make(map[string]string),
		timer: make(map[string]string),
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
	cache.timer[key] = deadline.String()
}

func (cache Cache) CheckTimer() {
	for key := range cache.storage {
		if timeString, ok := cache.timer[key]; ok {
			dueDate, error := time.Parse("dd-mm-yyyy HH:mm:ss", timeString)
			if error != nil {
				fmt.Println(error)
				return
			}

			if dueDate.Before(time.Now()) {
				delete(cache.storage, key)
				delete(cache.timer, key)
			}
		}
	}
}
