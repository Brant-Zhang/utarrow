package main

import ()

type avlTree struct {
	root *treeNode
}

func (a *avlTree) init() *avlTree {
	a.root = nil
	return a
}

func NewTree() *avlTree {
	new(avlTree).init()
}

func (a *avlTree) Insert(v elementType) {
	if a.root == nil {
		r := newNode(v)
		a.root = r
	}
}
