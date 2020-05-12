package main

type timer struct {
	i     int
	score int64
	key   string
}

type timerst struct {
	t []*timer
}

func newTimer() *timerst {
	t := new(timerst)
	return t
}

func (s *timerst) addSession(m *timer) {
	m.i = len(s.t)
	s.t = append(s.t, m)
	s.shiftUp(m.i)
	if m.i == 0 {

	}
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
			w3 := t[c3].when
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

func main() {
	t := newTimer()
}
