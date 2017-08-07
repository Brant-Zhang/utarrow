package pattern

import "testing"

func TestSingle(t *testing.T) {
	s := NewSingle()
	s["this"] = "that"
	s2 := NewSingle()
	if s2["this"] != "that" {
		t.Fatal("singleton failed")
	}

}
