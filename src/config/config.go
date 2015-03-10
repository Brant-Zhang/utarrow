package config

import (
	"bufio"
	//"fmt"
	"io"
	"os"
	"strings"
)

type section map[string]string

func Readfile(path string, flag string) (section, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return readbuf(f, flag)
}

func readbuf(f io.Reader, flag string) (section, error) {
	s := make(section)
	r := bufio.NewReader(f)
	var err error
	var line string
	var getf bool = false
	for err == nil {
		line, err = r.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" || line[0] == '#' {
			continue
		}
		if line[0] == '[' && line[len(line)-1] == ']' {
			if string(line[1:len(line)-1]) == flag {
				getf = true
				continue
			} else {
				getf = false
			}
		}
		if getf == true {
			pairs := strings.SplitN(line, "=", 2)
			if len(pairs) != 2 {
				continue
			} else {
				s[pairs[0]] = pairs[1]
			}
		}

	}
	return s, nil
}

/*
func main() {
	var filePath = "./udpserv.ini"
	var section = "haha"
	Param, err := Readfile(filePath, section)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(Param)
}
*/
