package model

import "errors"

// Used by BigoObject Type attribute
const (
	BigoString    = "string"
	BigoList      = "list"
	BigoHashTable = "hashtable"
	BigoSet       = "set"

	// Key operations
	BigoKey = "key"
)

// Encoding constant
const (
	BigoEncodingInt       = 10
	BigoEncodingString    = 11
	BigoEncodingHashTable = 12
	BigoEncodingList      = 13
	BigoEncodingSet       = 14
)

var (
	commandNotSupported = errors.New("command not supported")
)
