package main

import (
	"fmt"
	"sync"
)

type timer struct {
	i     int
	score int64
	key   string
}

type timerst struct {
	t      []*timer
	tcap   int
	tindex int
	lock   sync.RWMutex
}

func newTimer() *timerst {
	t := new(timerst)
	t.tcap = 10
	t.t = make([]*timer, 10)
	return t
}

func (s *timerst) addSession(m *timer) {
	s.lock.Lock()
	defer s.lock.Unlock()
	m.i = len(s.t)
	s.t = append(s.t, m)
	s.shiftUp(m.i)
	if m.i == 0 {

	}
}

func (s *timerst) delSession(m *timer) bool {
	s.lock.Lock()
	i := m.i
	last := len(s.t) - 1
	if i < 0 || i > last || s.t[i] != m {
		s.lock.Unlock()
		return false
	}
	if i != last {
		s.t[i] = s.t[last]
		s.t[i].i = i
	}
	s.t[last] = nil
	s.t = s.t[:last]
	if i != last {
		s.shiftUp(i)
		s.shiftDown(i)
	}
	s.lock.Unlock()
	return true
}

func (s *timerst) shiftDown(i int) {
	t := s.t
	n := len(t)
	score := t[i].score
	tmp := t[i]
	for {
		c := i*4 + 1 // left child
		c3 := c + 2  // mid child
		if c >= n {
			break
		}
		w := t[c].score
		if c+1 < n && t[c+1].score < w {
			w = t[c+1].score
			c++
		}
		if c3 < n {
			w3 := t[c3].score
			if c3+1 < n && t[c3+1].score < w3 {
				w3 = t[c3+1].score
				c3++
			}
			if w3 < w {
				w = w3
				c = c3
			}
		}
		if w >= score {
			break
		}
		t[i] = t[c]
		t[i].i = i
		t[c] = tmp
		t[c].i = c
		i = c
	}
}

func (s *timerst) shiftUp(i int) {
	t := s.t
	score := t[i].score
	tmp := t[i]
	for i > 0 {
		p := (i - 1) / 4 // parent
		if score >= t[p].score {
			break
		}
		t[i] = t[p]
		t[i].i = i
		t[p] = tmp
		t[p].i = p
		i = p
	}
}

func (s *timerst) show() {
	for _, v := range s.t {
		fmt.Printf("score:%d,key:%s\n", v.score, v.key)
	}
}

func main() {
	t := newTimer()
	for i := 11; i < 15; i++ {
		m := new(timer)
		m.score = int64(i)
		m.key = "hello"
		t.addSession(m)
	}
	t.show()
}
