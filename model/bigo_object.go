package model

// a bigoObject represents a value of bigo key-value peer
type BigoValue struct {
	Type     string      `json:"type"` // Type can be BigoString, BigoList, BigoHashTable, BigoSet
	Encoding uint8       `json:"encoding"`
	Data     interface{} `json:"ptr"`
}
