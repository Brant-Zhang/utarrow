package zencrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	//"fmt"
)

type Block cipher.Block

type BlockMode cipher.BlockMode

type ecb struct {
	b         Block
	blockSize int
}

func newECB(b Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b Block) BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b Block) BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int {
	return x.blockSize
}

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

func pkcs7Padding(cipherText []byte) []byte {
	padding := (16 - len(cipherText)%16) % 16
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padtext...)
}

func pkcs7Unpadding(originData []byte) []byte {
	if len(originData) <= 0 {
		return []byte("")
	}
	bytelen := len(originData)
	unpadding := int(originData[bytelen-1])
	if unpadding > 15 {
		return originData
	}
	for i := 1; i <= unpadding; i++ {
		if originData[bytelen-i] != byte(unpadding) {
			return originData
		}
	}
	//fmt.Printf("%d====%d=\n", len(originData), unpadding)
	return originData[:bytelen-unpadding]
}

func Encrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}
	plainText := pkcs7Padding(origData)
	mode := NewECBEncrypter(block)
	ciphertext := make([]byte, len(plainText))
	mode.CryptBlocks(ciphertext, plainText)
	return ciphertext, nil
}

func Decrypt(cipher, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}
	if len(cipher)%16 != 0 {
		return []byte(""), errors.New("not 16")
	}
	mode := NewECBDecrypter(block)
	origData := make([]byte, len(cipher))
	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(origData, cipher)
	//fmt.Printf("2222:%v\n", origData)
	origData = pkcs7Unpadding(origData)
	return origData, nil
}
