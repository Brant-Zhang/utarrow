package pattern

import (
	"testing"
)

func TestDuck(t *testing.T) {
	d := NewMarkDuck()
	d.Display()
	d.PerformFly()
	d.PerformQuack()
}
