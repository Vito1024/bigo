package model

type Key struct {
	*baseType
}

func NewKey() *Key {
	key := &Key{
		newbaseType(),
	}

	return key
}

