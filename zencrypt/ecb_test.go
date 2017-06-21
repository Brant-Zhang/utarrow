package zencrypt

import (
	"testing"
)

var key = []byte("sdf128j08dfja274lsdhg53js0hr3a2b")

var srcData = []byte("hello world")

func TestEncrypt(t *testing.T) {
	c, err := Encrypt(srcData, key)
	if err != nil {
		t.Fatal(err)
	}
	o, err := Decrypt(c, key)
	if err != nil {
		t.Fatal(err)
	}
	if string(srcData) != string(o) {
		t.Fatal("failed decrypt the cipher data")
	}
	t.Log("ok")
}
