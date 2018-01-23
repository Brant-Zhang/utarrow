package shell

import (
	"testing"
)

func TestShell(t *testing.T) {
	v, err := ShellRun("ip rule list")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}
