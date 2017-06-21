package zencrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rc4"
	"crypto/sha1"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"fmt"
)

func AesEncryptPkcs7(origData, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}
	//blockSize := block.BlockSize()
	origData = pkcs7Padding(origData)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecryptPkcs7(crypted, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}
	if len(crypted)%16 != 0 {
		return []byte(""), errors.New("cipher len not 16s")
	}

	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData, err = pkcs7UnPadding(origData)
	if err != nil {
		return nil, errors.New("unpadding error")
	}
	return origData, nil
}

func AesEncryptZero(origData, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}
	//blockSize := block.BlockSize()
	origData = zeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	//fmt.Println(crypted)
	return crypted, nil
}

func AesDecryptZero(crypted, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}
	if len(crypted)%16 != 0 {
		return []byte(""), errors.New("cipher len not 16s")
	}

	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = zeroUnPadding(origData)
	return origData, nil
}

func zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize) % blockSize
	//padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	padtext := bytes.Repeat([]byte{byte(0)}, padding)
	return append(ciphertext, padtext...)
}

func zeroUnPadding(origData []byte) []byte {
	bytelen := len(origData)
	if bytelen <= 0 {
		return []byte("")
	}

	for {
		if origData[bytelen-1] != 0 || (bytelen == 0) {
			break
		} else {
			bytelen--
		}
	}
	return origData[:bytelen]
}

func pkcs7UnPadding(originData []byte) ([]byte, error) {
	//var err = errors.New("unpadding error")
	if len(originData) < 16 {
		return nil, errors.New("input length short than 16")
	}
	bytelen := len(originData)
	unpadding := int(originData[bytelen-1])
	if unpadding > 16 || unpadding < 1 {
		return nil, fmt.Errorf("unpadding:%x > bytelen:%x", unpadding, bytelen)
	}
	for i := bytelen - unpadding; i < bytelen; i++ {
		if originData[i] != byte(unpadding) {
			return nil, errors.New("unpadding not equal byte")
		}
	}
	return originData[:bytelen-unpadding], nil
}

func pkcs7Padding(cipherText []byte) []byte {
	padding := 16 - (len(cipherText) % 16)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padtext...)
}

func Md5Cal(data []byte) string {
	f := md5.New()
	f.Write(data)
	md5str := hex.EncodeToString(f.Sum(nil))
	return md5str
}

func Md5Cal2Byte(data []byte) []byte {
	f := md5.New()
	f.Write(data)
	return f.Sum(nil)
}

func ValidateMd5(data []byte, sum []byte) bool {
	mysum := Md5Cal2Byte(data)
	return bytes.Equal(mysum, sum)
}

func Sha1Cal(data []byte) string {
	fsha1 := sha1.New()
	fsha1.Write(data)
	fileh := hex.EncodeToString(fsha1.Sum(nil))
	return fileh
}

func Rc4Crypt(data []byte, key []byte) []byte {
	rc, _ := rc4.NewCipher(key)
	dst := make([]byte, len(data))
	rc.XORKeyStream(dst, data)
	return dst
}

func GetSum(buf []byte) uint16 {
	var a uint16 = 0xbeaf
	for _, v := range buf {
		a += uint16(v)
	}
	return a
}

func GobEncode(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func GobDecode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
