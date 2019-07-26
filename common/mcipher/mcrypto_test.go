package mcipher

import (
	"strings"
	"testing"
)

var (
	aeskey = []byte{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 11, 12, 13, 14, 15, 16}
	aesiv  = []byte{32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17}
	src    = "qweuioujaljsdkh[]p;/1023"
)

func TestAesEncrypPkcs7(t *testing.T) {
	cip, err := AesEncryptPkcs7([]byte(src), aeskey, aesiv)
	if err != nil {
		t.Fatal(err)
	}
	srcAfter, err := AesDecryptPkcs7([]byte(cip), aeskey, aesiv)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.EqualFold(src, string(srcAfter)) {
		t.Fatal("aes cipher auth failed")
	}
}

func TestRc4(t *testing.T) {
	srcAfter := Rc4Crypt([]byte(src), aeskey)
	srcRaw := Rc4Crypt(srcAfter, aeskey)
	if !strings.EqualFold(src, string(srcRaw)) {
		t.Fatal("rc4  failed")
	}
}

func TestGob(t *testing.T) {
	var data = make(map[string]interface{})
	data["name"] = "lucas"
	data["age"] = 32
	dc, err := GobEncode(data)
	if err != nil {
		t.Fatal(err)
	}

	var data2 = make(map[string]interface{})
	err = GobDecode(dc, &data2)
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range data {
		if v != data2[k] {
			t.Fatal("gob encode failed")
		}
	}
}
