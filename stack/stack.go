//基于链表实现
package stack

import (
	"errors"
	"fmt"
)

// Element is an element of a linked list.
type Element struct {
	next *Element
	// The value stored with this element.
	Value interface{}
}

// List represents a single-linked list.
// The zero value for List is an empty list ready to use.
type List struct {
	root Element // sentinel list element, only &root and root.next are used
	len  int     // current list length excluding (this) sentinel element
}

// Init initializes or clears list l.
func (l *List) Init() *List {
	l.root.next = nil
	l.len = 0
	return l
}

// New returns an initialized list.
func New() *List { return new(List).Init() }

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *List) Len() int { return l.len }

// lazyInit lazily initializes a zero List value.
func (l *List) lazyInit() {
	if l == nil {
		l.Init()
	}
}

// insert inserts e after at, increments l.len, and returns e.
func (l *List) insert(e, at *Element) *Element {
	n := at.next
	at.next = e
	e.next = n
	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// Push inserts a new element e with value v at the front of list l and returns e.
func (l *List) Push(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

func (l *List) Pop() interface{} {
	e := l.root.next
	if e == nil {
		return nil
	}
	l.root.next = e.next
	e.next = nil
	l.len--
	return e.Value
}

func (s *List) IsEmpty() bool {
	if s.root.next == nil {
		return true
	}
	return false
}

func (s *List) MakeEmpty() error {
	if s != nil {
		return errors.New("first create stack!")
	}

	for !s.IsEmpty() {
		s.Pop()
	}
	return nil
}

/*
func main() {
	s := New()
	s.Push(100)
	s.Push("hahahha")
	s.Push(1111)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}*/
