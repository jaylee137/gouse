package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/thuongtruong1009/gouse/types"
)

type Tmp struct {
	Expires time.Duration
	Set     *cache.Cache
}

func NewTmp(expires ...time.Duration) *Tmp {
	var expire time.Duration
	if len(expires) > 0 {
		expire = expires[0]
	} else {
		expire = 24 * time.Hour
	}
	return &Tmp{
		Expires: expire,
		Set:     cache.New(expire, expire),
	}
}

func (c *Tmp) GetTmp(cacheKey string) interface{} {
	if val, found := c.Set.Get(cacheKey); found {
		return val
	}
	return nil
}

func (c *Tmp) SetTmp(cacheKey string, value interface{}, expireTime time.Duration) {
	if expireTime == 0 {
		expireTime = cache.DefaultExpiration
	}

	c.Set.Set(cacheKey, value, expireTime)
}

func (c *Tmp) DelTmp(cacheKey string) {
	c.Set.Delete(cacheKey)
}

func (c *Tmp) FlushTmp() {
	c.Set.Flush()
}

func (c *Tmp) AllTmp() map[string]string {
	result := make(map[string]string)
	for k := range c.Set.Items() {
		result[k] = types.InterfaceToString(c.GetTmp(k))
	}
	return result
}
