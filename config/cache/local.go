package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var caching = cache.New(24*time.Hour, 24*time.Hour)

func GetLocal(cacheKey string) (i interface{}) {
	if val, found := caching.Get(cacheKey); found {
		return val
	}
	return nil
}

func SetLocal(cacheKey string, value interface{}, expireTime time.Duration) {
	if expireTime == 0 {
		expireTime = cache.DefaultExpiration
	}

	caching.Set(cacheKey, value, expireTime)
}

func DelLocal(cacheKey string) {
	caching.Delete(cacheKey)
}

func FlushLocal() {
	caching.Flush()
}

func AllLocal() map[string]cache.Item {
	return caching.Items()
}
