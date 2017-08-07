package pattern

import (
	"testing"
)

func TestLogging(t *testing.T) {
	n := NewNSQLookupd("info")
	for i := 0; i < 5; i++ {
		n.logf(i, "Test")
	}
}

func TestDecorator(t *testing.T) {
	f := LogDecorate(Double)
	r := f(5)
	t.Log(r)
}
