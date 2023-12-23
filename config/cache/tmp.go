package cache

import (
	"errors"
	"sync"
)

var set = make(map[string]string)
var lock sync.RWMutex

func GetTmp(key string) (string, error) {
	lock.RLock()
	defer lock.RUnlock()
	if set == nil {
		return "", errors.New("set cache is not initialized")
	}
	return set[key], nil
	return "", nil
}

func SetTmp(key, value string) {
	lock.Lock()
	defer lock.Unlock()
	set[key] = value
}

func DelTmp(key string) {
	lock.Lock()
	defer lock.Unlock()
	delete(set, key)
}

func FlushTmp() {
	lock.Lock()
	defer lock.Unlock()
	set = map[string]string{}
}

func AllTmp() map[string]string {
	lock.RLock()
	defer lock.RUnlock()
	return set
}
