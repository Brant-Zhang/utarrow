package list

import (
	"math/rand"
)

const (
	MAXLEVEL = 32
	SP       = 0.25
)

type SEelement struct {
	Value interface{}
	Key   int
	Level int
	next  []*SEelement
}

func randomLV() int {
	lv := 1
	for v := rand.Intn(4); v == 3; {
		lv += 1
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
	l.level = 0
	l.next
	return l
}

func NewSkipList() *SList {
	return new(SList).Init()
}

func (l *SList) Len() int { return l.len }

func (l *SList) Search(key int) *SEelement {
	var tmp *SEelement = l.root.next[l.level]
	var prev = make([]*SEelement, l.level)
	for lv := l.level; lv >= 0; lv-- {
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
	var tmp *SEelement = l.root.next[l.level]
	var prev = make([]*SEelement, l.level)
	for lv := l.level; lv >= 0; lv-- {
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
					break
				}
			}
		}
	}
	// new node
	e := newSElement(key, value)
	for lv := l.level; lv >= 0; lv-- {
		e.next[lv] = prev[lv].next[lv]
		prev[lv].next[lv] = e
	}
	return e
}
