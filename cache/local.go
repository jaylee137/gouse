package cache

import (
	"errors"
	"sync"
)

type Local struct {
	Set  map[string]string
	Lock sync.RWMutex
}

func NewLocal() *Local {
	return &Local{
		Set:  make(map[string]string),
		Lock: sync.RWMutex{},
	}
}

func (c *Local) GetLocal(key string) (string, error) {
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	if c.Set == nil {
		return "", errors.New("set cache is not initialized")
	}
	return c.Set[key], nil
}

func (c *Local) SetLocal(key, value string) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.Set[key] = value
}

func (c *Local) DelLocal(key string) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	delete(c.Set, key)
}

func (c *Local) FlushLocal() {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.Set = map[string]string{}
}

func (c *Local) AllLocal() map[string]string {
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	return c.Set
}
