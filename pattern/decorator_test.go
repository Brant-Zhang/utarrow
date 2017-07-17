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
