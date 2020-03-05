package list

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MAXLEVEL = 32
	SP       = 0.25
)

var RD *rand.Rand

func pre() {
	s := rand.NewSource(time.Now().Unix())
	RD = rand.New(s)
}

type SEelement struct {
	Value interface{}
	Key   int
	Level int
	next  []*SEelement
}

func randomLV() int {
	lv := 1
	for {
		if v := RD.Intn(4); v == 3 {
			lv += 1
		} else {
			break
		}
	}
	return lv
}

func newSElement(key int, v interface{}) *SEelement {
	e := new(SEelement)
	e.Key = key
	e.Value = v
	e.Level = randomLV()
	e.next = make([]*SEelement, e.Level)
	return e
}

func (e *SEelement) Next(l int) *SEelement {
	if e.Level < l {
		return nil
	}
	return e.next[l]
}

type SList struct {
	root  SEelement
	len   int
	level int
}

func (l *SList) Init() *SList {
	l.len = 0
	l.level = 1
	l.root.next = make([]*SEelement, 1)
	return l
}

func NewSkipList() *SList {
	pre()
	return new(SList).Init()
}

func (l *SList) Len() int { return l.len }

func (l *SList) Search(key int) *SEelement {
	var tmp *SEelement = l.root.next[l.level-1]
	var prev = make([]*SEelement, l.level)
	for lv := l.level - 1; lv >= 0; lv-- {
		if tmp != nil {
			continue
		}
		for tmp.next != nil {
			if tmp.Key < key {
				prev[lv] = tmp
				tmp = tmp.next[lv]
			} else if tmp.Key == key {
				return tmp
			} else {
				//position>target;next level
				continue
			}
		}
		//tmp.next=nil
	}
	return nil
}

func (l *SList) Push(key int, value interface{}) *SEelement {
	var tmp *SEelement = l.root.next[l.level-1]
	var prev = make([]*SEelement, l.level)
lb:
	for lv := l.level - 1; lv >= 0; lv-- {
		prev[lv] = &l.root
		for tmp != nil {
			if tmp.Key < key {
				prev[lv] = tmp
				tmp = tmp.next[lv]
			} else if tmp.Key == key {
				//update
				tmp.Value = value
				return tmp
			} else {
				if lv == 0 {
					break lb
				}
			}
		}
	}
	// new node
	e := newSElement(key, value)
	var lg int
	if e.Level > l.level {
		lg = l.level
	} else {
		lg = e.Level
	}
	for lv := lg - 1; lv >= 0; lv-- {
		fmt.Println("--------", lv, l.level)
		e.next[lv] = prev[lv].next[lv]
		prev[lv].next[lv] = e
	}
	for e.Level > l.level {
		var add *SEelement = e
		l.root.next = append(l.root.next, add)
		l.level++
	}
	return e
}
