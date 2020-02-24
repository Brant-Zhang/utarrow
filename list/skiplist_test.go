package list

import (
	"testing"
)

func TestShow(t *testing.T) {
	l := NewSkipList()
	l.Push(3, "japan")
	l.Push(10, "china")
	l.Push(22, "taiwan")

	err := Setup("/tmp/log", "debug")
	if err != nil {
		t.Error("log error", err)
	}
	Info("good morning brant!")
	Debugln("hello 1")
	Warnln("hello 2")
	var a []byte
	var b int
	Debug("nihao--%s--%v", a, b)
}

/*
func TestXxx(t *testing.T){
	t.Error("no pass")
}
*/
