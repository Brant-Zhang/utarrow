package safeset

import (
	"testing"
	"time"
)

func TestGetsort(t *testing.T) {
	s := Newset(0, 3, 1)
	s.Add("zhang", 11)
	s.Add("li", 12)
	s.Add("wang", 13)
	s.Delete("li")
	en := s.GetSet()
	t.Log(en) //show all items
	time.Sleep(time.Second * 2)
	s.Add("zhang", 10) //update one item
	time.Sleep(time.Second * 3)
	en = s.GetSet()
	t.Log(en)
	if len(en) != 1 {
		t.Fatal("not fit")
	}

}
