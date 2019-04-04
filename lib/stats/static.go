package stats

import (
	"fmt"
	"sync/atomic"
	"time"
)

type statesVal struct {
	v int64
}

func Increment(sv *statesVal) {
	atomic.AddInt64(&sv.v, 1)
}

func Decrement(sv *statesVal) {
	atomic.AddInt64(&sv.v, -1)
}

func Reset(sv *statesVal){
	atomic.StoreInt64(&sv.v,0)
}

func Assign(sv *statesVal, count int64) {
	sv.v = count
}

func AddStateVal(sv *statesVal) {
	sv.v++
}

type statsM map[string]*statesVal

func NewStats() statsM{
	v:=make(statsM,0)	
	return v
}

func (s statsM)AddKey(k string){
	s[k]=&statesVal{0}
}

func (s statsM)PrintStats() {
	fmt.Printf("%s:", time.Now().Format(time.RFC3339))
	for k, v := range s {
		fmt.Printf("%s:%d;\t\t", k, v.v)
	}
	fmt.Printf("\n")
}

func (s statsM)showStatus() {
	t3 := time.NewTicker(time.Second * 120) //print statics
	for {
		select {
		case <-t3.C:
			s.PrintStats()
		}
	}
}
