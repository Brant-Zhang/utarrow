package ttlevent

import (
	"testing"
)

func Test_common(t *testing.T) {
	c := NewCache(30)
	c.Put("step1")
	c.Put("step2")
	if err := c.backup(); err != nil {
		t.Fatal(err)
	}
	t.Log("hello")
}

func Test_restore(t *testing.T) {
	c := NewCache(30)
	b := len(c.pool)
	if b == 0 {
		t.Fatal("restore failed")
	}
	t.Log("success")
}
