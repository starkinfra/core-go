package cache

import (
	"sync"
	"testing"
)

func TestCacheConcurrentAccess(t *testing.T) {
	key := "test-concurrent-access"
	t.Cleanup(func() {
		lock.Lock()
		defer lock.Unlock()
		delete(Cache, key)
	})
	var group sync.WaitGroup
	for i := 0; i < 100; i++ {
		group.Add(2)
		go func(value int) {
			defer group.Done()
			Set(key, value)
		}(i)
		go func() {
			defer group.Done()
			Get(key)
		}()
	}
	group.Wait()
	if Get(key) == nil {
		t.Errorf("expected cached value, got nil")
	}
}
