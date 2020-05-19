package datastructure

import (
	"sort"
)

type btnode struct {
	value int
	lchild *btnode
	rchild *btnode
}

type BinaryTree struct {
	root *btnode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{nil}
}

func (bt *BinaryTree) GetAll() (res []int) {
	bt.PreOrder(bt.root, &res)
	return res
}

func (bt *BinaryTree) PreOrder(n *btnode, res *[]int) {
	if n == nil {
		return
	}
	*res = append(*res, n.value)
	bt.PreOrder(n.lchild, res)
	bt.PreOrder(n.rchild, res)
}

func (bt *BinaryTree) Set(v ...int) {
    sort.Ints(v)
	if len(v) == 0 { return }
	q := queue{}
	bt.root = &btnode{v[0], nil, nil}
	v = v[1:]
	q.push(bt.root)
	for len(v) != 0 {
		nd := q.pop()
		nd.lchild = &btnode{v[0], nil, nil}
		v = v[1:]
		if len(v) == 0 { return }
		nd.rchild = &btnode{v[0], nil, nil}
		v = v[1:]
		q.push(nd.lchild)
		q.push(nd.rchild)
	}
}

type queue struct {
	values []*btnode
}

func (q *queue) push(v *btnode) {
	q.values = append(q.values, v)
}

func (q *queue) pop() *btnode {
	if len(q.values) == 0 {
		return nil
	}

	v := q.values[0]
	q.values = q.values[1:]
	return v
}
