package model

type Set struct {
	*baseType
}

func NewSet() *Set {
	return &Set{ newbaseType() }
}


