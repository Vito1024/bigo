package model

type HashTable struct {
	*baseType
}

func NewHashTable() *HashTable {
	return &HashTable{newbaseType()}
}
