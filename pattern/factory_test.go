package pattern

import (
	"io"
	"testing"
)

func TestWirte(t *testing.T) {
	s := NewStore(DiskStorage)
	f, err := s.Open("z.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	n, err := f.Write([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("data write counts:", n)
}

func TestRead(t *testing.T) {
	s := NewStore(DiskStorage)
	f, err := s.Open("z.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	var buf = make([]byte, 1024)
	n, err := f.Read(buf)
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}
	t.Log("data read:", string(buf[:n]))

}
