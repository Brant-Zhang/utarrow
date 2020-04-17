package shell

import (
	"testing"
)

func TestShell(t *testing.T) {
	v, err := ShellRun("systemctl start redis")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}
