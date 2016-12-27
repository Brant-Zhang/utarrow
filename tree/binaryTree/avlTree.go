//带有平衡条件的二叉查找树，保证树的深度为O(logn)，每个结点左右子树高度最多相差1(平衡性)，也就保证了查找时间复杂度为O(logn)
//对某个结点的插入或删除可能改变树的平衡性，此时需要旋转恢复平衡
/*
假设有一个结点的平衡因子为2（在AVL树中，最大就是2，因为结点是一个一个地插入到树中的，一旦出现不平衡的状态就会立即进行调整，因此平衡因子最大不可能超过2），那么就需要进行调整。由于任意一个结点最多只有两个儿子，所以当高度不平衡时，只可能是以下四种情况造成的：
1. 对该结点的左儿子的左子树进行了一次插入。
2. 对该结点的左儿子的右子树进行了一次插入。
3. 对该结点的右儿子的左子树进行了一次插入。
4. 对该结点的右儿子的右子树进行了一次插入。
情况1和4是关于该点的镜像对称，同样，情况2和3也是一对镜像对称。因此，理论上只有两种情况
*/
//references:
//https://en.wikipedia.org/wiki/AVL_tree
//
package main

import (
	"fmt"
)

type elementType int

type avlTree struct {
	element elementType
	lchild  *avlTree
	rchild  *avlTree
	height  int
}

func (t *avlTree) Height() int {
	if t == nil {
		return -1
	}
	return t.height
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Newtree(v elementType) *avlTree {
	return &avlTree{
		element: v,
		lchild:  nil,
		rchild:  nil,
		height:  0,
	}
}

//can be called only if t has a left child
func (t *avlTree) SingleRotateWithLeft() *avlTree {
	var tp *avlTree
	tp = t.lchild
	t.lchild = tp.rchild
	tp.rchild = t
	t.height = Max(t.lchild.Height(), t.rchild.Height()) + 1
	tp.height = Max(tp.lchild.Height(), tp.rchild.Height()) + 1
	return tp
}

//can be called only if t has a right child
func (t *avlTree) SingleRotateWithRitht() *avlTree {
	var tp *avlTree
	tp = t.rchild
	tp.lchild = t
	t.rchild = tp.lchild
	t.height = Max(t.lchild.Height(), t.rchild.Height()) + 1
	tp.height = Max(tp.lchild.Height(), tp.rchild.Height()) + 1
	return tp
}

//can be called only if t has a left child and its left child has a right child
func (t *avlTree) DoubleRotateWithLeft() *avlTree {
	t.lchild = t.SingleRotateWithRitht()
	return t.SingleRotateWithLeft()
}

func (t *avlTree) DoubelRotateWithRight() *avlTree {
	t.rchild = t.SingleRotateWithLeft()
	return t.SingleRotateWithRitht()
}

func (t *avlTree) Insert(v elementType) *avlTree {
	if t == nil {
		t = &avlTree{
			element: v,
			lchild:  nil,
			rchild:  nil,
			height:  0,
		}
		return t
	}
	if v < t.element {
		t.lchild = t.lchild.Insert(v)
		if t.lchild.Height()-t.rchild.Height() == 2 {
			if v < t.lchild.element {
				t = t.SingleRotateWithLeft()
			} else {
				t = t.DoubleRotateWithLeft()
			}
		}
	} else if v > t.element {
		t.rchild = t.rchild.Insert(v)
		if t.rchild.Height()-t.lchild.Height() == 2 {
			if v > t.rchild.element {
				t = t.SingleRotateWithRitht()
			} else {
				t = t.DoubleRotateWithLeft()
			}
		}
	} //else v=t.element, do nothing

	t.height = Max(t.lchild.Height(), t.rchild.Height()) + 1
	return t
}

//找到该节点后用其右子树的左子树替代该节点
func (t *avlTree) Delete(v elementType) *avlTree {
	if t == nil {
		return nil
	}
	if t.element == v {
		//右子树不存在，用左子树替代
		if t.rchild == nil {
			//tp := t
			t = t.lchild
			//tp = nil
		} else {
			tp := t.rchild
			for tp != nil {
				tp = tp.lchild
			}
			t.element = tp.element
			t.rchild = t.rchild.Delete(t.element)
			t.height = Max(t.lchild.Height(), t.rchild.Height()) + 1
		}
		return t
	} else if v < t.element {
		t.lchild = t.lchild.Delete(v)
		if t.lchild != nil {
			t.lchild = t.lchild.Rotate()
		}
	} else {
		t.rchild = t.rchild.Delete(v)
		if t.rchild != nil {
			t.rchild = t.rchild.Rotate()
		}
	}
	t = t.Rotate()
	return t
}

func (t *avlTree) Rotate() *avlTree {
	if t.lchild.Height()-t.rchild.Height() == 2 {
		if t.lchild.lchild.Height() > t.lchild.rchild.Height() {
			t = t.SingleRotateWithLeft()
		} else {
			t = t.DoubleRotateWithLeft()
		}
	} else if t.rchild.Height()-t.lchild.Height() == 2 {
		if t.rchild.lchild.Height() > t.rchild.rchild.Height() {
			t = t.DoubelRotateWithRight()
		} else {
			t = t.SingleRotateWithRitht()
		}
	}
	return t
}

func (t *avlTree) MidOrderTraversal() {
	if t != nil {
		t.lchild.MidOrderTraversal()
		fmt.Println(t.element)
		t.rchild.MidOrderTraversal()
	}
}

func main() {
	t := Newtree(10)
	t.Insert(8)
	t.Insert(24)
	t.Insert(13)
	t.Insert(99)
	t.Insert(15)
	t.Insert(3)
	t.Insert(4)
	t.Insert(1)
	t.MidOrderTraversal()
}
