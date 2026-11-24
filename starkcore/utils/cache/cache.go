package cache

import "sync"

// Deprecated: read and write through Get and Set; direct map access is not safe for concurrent use.
var Cache = map[string]interface{}{}

var lock sync.RWMutex

func Get(key string) interface{} {
	lock.RLock()
	defer lock.RUnlock()
	return Cache[key]
}

func Set(key string, value interface{}) {
	lock.Lock()
	defer lock.Unlock()
	Cache[key] = value
}
