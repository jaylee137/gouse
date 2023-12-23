package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/thuongtruong1009/gouse/types"
)

var caching = cache.New(24*time.Hour, 24*time.Hour)

func GetTmp(cacheKey string) interface{} {
	if val, found := caching.Get(cacheKey); found {
		return val
	}
	return nil
}

func SetTmp(cacheKey string, value interface{}, expireTime time.Duration) {
	if expireTime == 0 {
		expireTime = cache.DefaultExpiration
	}

	caching.Set(cacheKey, value, expireTime)
}

func DelTmp(cacheKey string) {
	caching.Delete(cacheKey)
}

func FlushTmp() {
	caching.Flush()
}

func AllTmp() map[string]string {
	result := make(map[string]string)
	for k := range caching.Items() {
		result[k] = types.InterfaceToString(GetTmp(k))
	}
	return result
}
