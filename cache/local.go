package cache

import (
	"errors"
	"sync"
)

var set = make(map[string]string)
var lock sync.RWMutex

func GetLocal(key string) (string, error) {
	lock.RLock()
	defer lock.RUnlock()
	if set == nil {
		return "", errors.New("set cache is not initialized")
	}
	return set[key], nil
}

func SetLocal(key, value string) {
	lock.Lock()
	defer lock.Unlock()
	set[key] = value
}

func DelLocal(key string) {
	lock.Lock()
	defer lock.Unlock()
	delete(set, key)
}

func FlushLocal() {
	lock.Lock()
	defer lock.Unlock()
	set = map[string]string{}
}

func AllLocal() map[string]string {
	lock.RLock()
	defer lock.RUnlock()
	return set
}
