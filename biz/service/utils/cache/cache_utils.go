package cache

import (
	"sync"
	"time"
)

const (
	// DefaultExpirationTime is the default timeout for cache items, set to 1600 seconds
	DefaultExpirationTime = 1600 * time.Second
)

var (
	cacheMu sync.Mutex
	cache   = make(map[string]cacheItem)
)

type cacheItem struct {
	Value      string
	Expiration time.Time
}

// Get retrieves a value from the cache by its key.
func Get(key string) (string, bool) {
	cacheMu.Lock()
	defer cacheMu.Unlock()

	item, found := cache[key]
	if !found {
		return "", false
	}

	// Check if the item has expired
	if time.Now().After(item.Expiration) {
		delete(cache, key)
		return "", false
	}

	return item.Value, true
}

// Set adds or updates a value in the cache with a default expiration duration.
func Set(key string, value string) {
	cacheMu.Lock()
	defer cacheMu.Unlock()

	expirationTime := time.Now().Add(DefaultExpirationTime)

	cache[key] = cacheItem{
		Value:      value,
		Expiration: expirationTime,
	}
}

// Delete removes a value from the cache by its key.
func Delete(key string) {
	cacheMu.Lock()
	defer cacheMu.Unlock()

	delete(cache, key)
}

// Set adds or updates a value in the cache with an optional expiration duration.
func SetWithTimeout(key string, value string, expiration time.Duration) {
	cacheMu.Lock()
	defer cacheMu.Unlock()

	var expirationTime time.Time
	if expiration > 0 {
		expirationTime = time.Now().Add(expiration)
	} else {
		// Set a default expiration duration here (e.g., 30 minutes)
		expirationTime = time.Now().Add(30 * time.Minute)
	}

	cache[key] = cacheItem{
		Value:      value,
		Expiration: expirationTime,
	}
}
