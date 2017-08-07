package pattern

import "testing"

func TestOpenClose(t *testing.T) {
	c := &Circle{
		diameter: 10,
	}
	g := new(GrapheEditor)
	g.drawShape(c)
}
