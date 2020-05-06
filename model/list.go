package model

type List struct {
	*baseType
}

func NewList() *List {
	return &List{newbaseType()}
}
