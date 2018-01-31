package shell

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"
)

func ShellRun(line string) (string, error) {
	shell := os.Getenv("SHELL")
	b, err := exec.Command(shell, "-c", line).Output()
	if err != nil {
		if eerr, ok := err.(*exec.ExitError); ok {
			b = eerr.Stderr
		}
		return "", errors.New(err.Error() + ":" + string(b))
	}
	return strings.TrimSpace(string(b)), nil
}
func CmdRun(bin string, args ...string) (data []byte, err error) {
	cmd := exec.Command(bin, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return
	}
	data = out.Bytes()
	return
}
