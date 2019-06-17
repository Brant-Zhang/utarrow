package stats

import (
	"fmt"
	"sync/atomic"
	"time"
)

type statesVal struct {
	V int64
}

func Increment(sv *statesVal) {
	atomic.AddInt64(&sv.V, 1)
}

func IncrementNums(sv *statesVal, cn int) {
	atomic.AddInt64(&sv.V, int64(cn))
}

func Decrement(sv *statesVal) {
	atomic.AddInt64(&sv.V, -1)
}

func Reset(sv *statesVal) {
	atomic.StoreInt64(&sv.V, 0)
}

func Assign(sv *statesVal, count int64) {
	sv.V = count
}

func AddStateVal(sv *statesVal) {
	sv.V++
}

type StatsM map[string]*statesVal

func NewStats() StatsM {
	v := make(StatsM, 0)
	return v
}

func (s StatsM) AddKey(k string) {
	s[k] = &statesVal{0}
}

func (s StatsM) PrintStats() {
	fmt.Printf("%s:", time.Now().Format(time.RFC3339))
	for k, v := range s {
		fmt.Printf("%s:%d;\t\t", k, v.V)
	}
	fmt.Printf("\n")
}

func (s StatsM) Reset() {
	for _, v := range s {
		atomic.StoreInt64(&v.V, 0)
	}
}

func (s StatsM) showStatus() {
	t3 := time.NewTicker(time.Second * 120) //print statics
	for {
		select {
		case <-t3.C:
			s.PrintStats()
		}
	}
}
