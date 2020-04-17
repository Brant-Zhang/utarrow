package list

import (
	"testing"
)

func TestShow(t *testing.T) {
	l := NewSkipList()
	l.Push(3, "japan")
	l.Push(10, "china")
	l.Push(22, "taiwan")
	e := l.Search(22)
	if e != nil {
		t.Log(e.Value)
	}
}

/*
func TestXxx(t *testing.T){
	t.Error("no pass")
}
*/
