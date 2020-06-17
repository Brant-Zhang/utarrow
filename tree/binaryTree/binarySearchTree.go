//二叉树：每个节点不能有多于两个子节点
//下面实现一个二叉查找树（无重复值），它是一种有序的二叉树：对于每个结点，左子树（所有结点）小于自己，右子树（所有结点）大于自己
//如果插入的是个有序列，则此树退化成链表，操作的时间复杂度降为0(n)
package main

import (
	"fmt"
)

type elementType int
type position *treeNode
type treeNode struct {
	value  elementType
	lchild *treeNode
	rchild *treeNode
}

func (t *treeNode) Find(key elementType) position {
	if t == nil {
		return nil
	}
	if key < t.lchild.value {
		return t.lchild.Find(key)
	} else if key > t.rchild.value {
		return t.rchild.Find(key)
	} else {
		return t
	}
}

func (t *treeNode) LeftSearch() {
	if t == nil {
		return
	}
	fmt.Println(t.value)
	t.lchild.LeftSearch()
	t.rchild.LeftSearch()
}

func (t *treeNode) Insert(key elementType) *treeNode {
	if t == nil {
		t = new(treeNode)
		t.value = key
		t.lchild = nil
		t.rchild = nil
	}
	if key < t.value {
		t.lchild = t.lchild.Insert(key)
	} else if key > t.value {
		t.rchild = t.rchild.Insert(key)
	}
	return t
}

func (t *treeNode) FindMin() *treeNode {
	if t != nil {
		for t.lchild != nil {
			t = t.lchild
		}
	}

	return t
}

func (t *treeNode) FindMax() *treeNode {
	if t == nil {
		return nil
	}
	if t.rchild != nil {
		return t.rchild.FindMax()
	} else {
		return t
	}
}

//缺点：经过多次添加删除后，树会失去平衡，树的深度大大增加
//通常会用lazy detection优化，用删除标记代替实际删除
func (t *treeNode) Delete(key elementType) *treeNode {
	var tp *treeNode
	if t == nil {
		return nil
	}
	if key < t.value {
		t.lchild = t.lchild.Delete(key)
	} else if key > t.value {
		t.rchild = t.rchild.Delete(key)
	} else if t.lchild != nil && t.rchild != nil {
		tp = t.rchild.FindMin() //用右子树最小值代替被删除的值
		t.value = tp.value
		t.rchild = t.rchild.Delete(t.value)
	} else { //one or zero child
		tp = t
		if t.lchild == nil {
			t = t.rchild
		} else if t.rchild == nil {
			t = t.lchild
		}
		tp = nil
	}
	return t
}

func (t *treeNode) MakeEmpty() {
	if t != nil {
		t.lchild.MakeEmpty()
		t.rchild.MakeEmpty()
		t = nil
	}
}

func Newtree(v elementType) *treeNode {
	return &treeNode{
		value: v,
	}
}

func main() {
	bt := Newtree(10)
	bt.Insert(20)
	bt.Insert(1)
	bt.Insert(88)
	bt.Insert(100)
	bt.LeftSearch()
	//v := bt.Delete(10)
	//v = bt.FindMax()
	//if v != nil {
	//	fmt.Println(v.value)
	//}
	//fmt.Println(bt.value)
}
