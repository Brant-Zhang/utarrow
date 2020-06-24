package main

import "fmt"

type room struct {
	name string
	path string
	isps []string
	tws  []string
}

type Trie struct {
	kind     kind
	path     string
	parent   *Trie
	children []*Trie
	room     room
}

/** Initialize your data structure here. */
func Constructor() Trie {
	t := new(Trie)
	t.kind = bkind
	return *t
}

func (this *Trie) show() {
	fmt.Printf("path:%s\n", this.path)
	for _, v := range this.children {
		v.show()
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(r room) {
	word := r.path
	if word[0] != '/' {
		word = "/" + word
	}
	pl := len(this.path)
	wl := len(word)
	max := pl
	if wl < max {
		max = wl
	}
	//确定相同前缀位置
	l := 0
	for ; l < max && word[l] == this.path[l]; l++ {
	}
	//as root
	if l == 0 {
		this.path = word
		return
	} else if l < pl {
		//split this node
		t := new(Trie)
		t.path = this.path[l:pl]
		this.path = this.path[:l]
		this.children = append(this.children, t)
		if wl > l {
			nc := new(Trie)
			nc.path = word[l:wl]
			this.children = append(this.children, nc)
		}
	} else if l < wl {
		//add as children
		t := new(Trie)
		t.kind = leafkind
		t.path = word[l:wl]
		t.room = r
		this.children = append(this.children, t)
	} else {
		//node already exists
		//drop here
	}

}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if word[0] != '/' {
		word = "/" + word
	}
	for this != nil {
		pl := len(this.path)
		fmt.Printf("++++----------:%s\n", this.path)
		if pl < len(word) {
			if this.path == word[:pl] {
				for _, v := range this.children {
					fmt.Printf("++++:%s\n", word[pl:])
					if v.Search(word[pl:]) {
						return true
					}
				}
			}
		} else if pl > len(word) {
			return false
		} else if this.path == word {
			return true
		}
		return false
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return false
}
