package main

type elementType int

type treeNode struct {
	element elementType
	lchild  *treeNode
	rchild  *treeNode
	height  int
	sz      int //the number of nodes in subtree
}

func (t *treeNode) init(v elementType) *treeNode {
	t.lchild = nil
	t.rchild = nil
	t.element = v
	t.height = 0
	t.sz = 0
	return t
}

func newNode(v elementType) *treeNode {
	return new(treeNode).init(v)
}

//size of left subtree
func (t *treeNode) leftSize() int {
	if t.lchild == nil {
		return 0
	} else {
		return t.lchild.sz
	}
}

func (t *treeNode) rightSize() int {
	if t.rchild == nil {
		return 0
	} else {
		return t.rchild.sz
	}
}
