package datastructure

import "testing"

func TestBinaryTree_Set(t *testing.T) {
	bt := NewBinaryTree()
	bt.Set(10, 20, 3, 1, 18, 24, 2, 9, 12)
	t.Log(bt.GetAll())
}
