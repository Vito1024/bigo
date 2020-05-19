package model

type BinaryTree struct {
	*baseType
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{newbaseType()}
}
