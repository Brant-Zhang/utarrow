//线程安全的buffer缓冲池，可以被多goroutine同时获取／回收
package main

import (
	"errors"
	"fmt"
	"time"
)

var ERRBufPoolFull = errors.New("Gave a buffer to a full pool.")

const datasize = 1024
const BUFFERCOUNT = 1000000

type buffer struct {
	data [datasize]byte
	ulen int
}
type bchan chan *buffer

func newBuffer() bchan {
	b := make(bchan, BUFFERCOUNT)
	for i := 0; i < BUFFERCOUNT; i++ {
		buf := new(buffer)
		b.put(buf)
	}
	return b
}

/***************************************
//还可以设计成这种不需要初始填充的方式，
//并且缓冲池大小可以一直增加，不会阻塞
func (pool bchan) get() (buf *buffer) {
	select {
	case buf = <-pool:
	default:
		buf = new(buffer)
	}
	return buf
}
***************************************/

func (pool bchan) put(v *buffer) error {
	select {
	case pool <- v:
	default:
		return ERRBufPoolFull
	}
	return nil
}

func (pool bchan) get() *buffer {
	return <-pool
}

func main() {
	gpool := newBuffer()
	start := time.Now().UnixNano()
	for i := 0; i < BUFFERCOUNT; i++ {
		v := gpool.get()
		gpool.put(v)
	}
	end := time.Now().UnixNano()
	fmt.Println((end - start) / 1e6)
}
