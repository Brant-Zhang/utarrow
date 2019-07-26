package safeset

import (
	"fmt"
	"testing"
	"time"
)

func TestGetsort(t *testing.T) {
	s := Newset(0, 3, 1)
	s.Add("zhang", "china")
	s.Add("li", "us")
	s.Add("wang", "japan")
	s.Delete("li")
	en := s.GetSet()
	t.Log(en) //show all items
	time.Sleep(time.Second * 2)
	s.Add("zhang", "keroa") //update one item
	time.Sleep(time.Second * 3)
	en = s.GetSet()
	for _, v := range en {
		fmt.Println(v.v.(string))
	}

	if len(en) != 1 {
		t.Fatal("not fit")
	}

}

func TestPsort(t *testing.T) {
	s := Newpset()
	s.Add("zhang", "china")
	s.Add("li", "us")
	s.Add("wang", "japan")
	s.Delete("li")
	en := s.GetSet()
	t.Log(en) //show all items
	time.Sleep(time.Second * 2)
	s.Add("zhang", "keroa") //update one item
	time.Sleep(time.Second * 3)
	en = s.GetSet()
	for _, v := range en {
		t.Log(v.v.(string))
	}

	if len(en) <= 1 {
		t.Fatal("not fit")
	}

}
