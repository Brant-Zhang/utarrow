package ttlevent

import (
	"testing"
)

func remove(key string) {
	DefaultCache.mu.Lock()
	delete(DefaultCache.pool, key)
	DefaultCache.mu.Unlock()
}

func Test_common(t *testing.T) {
	c := NewCache(30)
	c.Put("step33")
	remove("step83")
	if err := c.backup(); err != nil {
		t.Fatal(err)
	}
	t.Log("hello")
}

func Test_restore(t *testing.T) {
	c := NewCache(30)
	b := len(c.pool)
	if b != 2 {
		t.Fatal("restore failed", b)
	}
	t.Log("success")
}
