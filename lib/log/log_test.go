package log

import (
	"testing"
)

func TestLogwrite(t *testing.T) {
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
