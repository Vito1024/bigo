package model

import "unsafe"

// a bigoObject represents a value of bigo key-value peer
type BigoObject struct {
	Type     uint8          `json:"type"` // Type can be BigoString, BigoList, BigoHashTable, BigoSet
	Encoding uint8          `json:"encoding"`
	Ptr      unsafe.Pointer `json:"ptr"`
}
