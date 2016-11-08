//ringbuffer 无锁环形队列是线程安全的（针对一个goroutine写多个goroutine读的情况），例如goroutine1正在对index=30的slot写，
//此时tail=29;goroutine2无法对index=30的slot进行读，只有数据写入完成才能被读到
package main

import (
	"fmt"
)

const bufsize = 128

type objectType interface{}

type ringbuf struct {
	head int
	tail int
	buf  []objectType
	sz   int
}

func newRingbuf(sz int) *ringbuf {
	return &ringbuf{
		head: 0,
		tail: 0,
		buf:  make([]objectType, sz),
		sz:   sz,
	}
}

func (r *ringbuf) empty() bool {
	return r.head == r.tail
}

func (r *ringbuf) full() bool {
	return (r.tail+1)%len(r.buf) == r.head
}

func (r *ringbuf) put(val objectType) bool {
	if r.full() {
		return false
	}
	r.buf[r.tail] = val
	r.tail = (r.tail + 1) % r.sz
	return true
}

func (r *ringbuf) get() (val objectType) {
	if r.empty() {
		return nil
	}
	val = r.buf[r.head]
	r.head = (r.head + 1) % r.sz
	return val
}

func (r *ringbuf) getall() (val []objectType) {
	if r.empty() {
		return nil
	}
	cptail := r.tail
	var l int
	if cptail > r.head {
		l = cptail - r.head
	} else {
		l = cptail - r.head + r.sz
	}

	val = make([]objectType, l)

	if cptail > r.head {
		copy(val, r.buf[r.head:cptail])
	} else {
		copy(val, r.buf[r.head:r.sz])
		copy(val[r.sz-r.head:], r.buf[:cptail])
	}
	return val
}

func main() {
	buf := newRingbuf(bufsize)
	var b bool
	for i := 0; i < 128; i++ {
		b = buf.put(i)
		if !b {
			fmt.Printf("buffer is full%d!\n", i)
			break
		}
	}

	var v interface{}
	for i := 0; i < 10; i++ {
		v = buf.get()
		if v == nil {
			break
		}
		//fmt.Println(v)
	}
	buf.put(999)
	vv := buf.getall()
	fmt.Println(vv)
}
