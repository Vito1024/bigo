package controller

import (
	"bigo/db"
	"bigo/model"
)

/*
	Abstract data type supported to BigoType interface
*/

var (
	Bigo = make(map[string]model.BigoType)

	Key = model.NewKey()

	String    = model.NewString()
	HashTable = model.NewHashTable()
	List      = model.NewList()
	Set       = model.NewSet()
)

func init() {

	Bigo[model.BigoKey] = Key
	//
	Bigo[model.BigoString] = String
	Bigo[model.BigoList] = List
	Bigo[model.BigoHashTable] = HashTable
	Bigo[model.BigoSet] = Set

	// bigoKey
	bigoKey()

	// bigoString
	bigoString()

	// bigoList
	bigoList()

	// bigoHashTable
	bigoHashTable()

	// bigoSet
	bigoSet()
}

// bigoKey namespace
func bigoKey() {
	Key.Register("DEL", db.KeyDEL)
	Key.Register("TYPE", db.KeyTYPE)
	Key.Register("KEY", db.KeyKEY)
	Key.Register("PING", db.KeyPING)
	Key.Register("SELECT", db.KeySELECT)
	Key.Register("DB", db.KeyDB)
	Key.Register("COUNT", db.KeyCOUNT)
	Key.Register("DUMP", db.KeyDUMP)
	Key.Register("RECOVER", db.KeyRECOVER)
}

// bigoString namespace
func bigoString() {
	String.Register("GET", db.StringGET)
	String.Register("SET", db.StringSET)
}

// bigoList namespace
func bigoList() {
	List.Register("LGET", db.ListGET)
	List.Register("LSET", db.ListSET)
	List.Register("APPEND", db.ListAPPEND)
	List.Register("LAPPEND", db.ListLAPPEND)
	List.Register("POP", db.ListPOP)
	List.Register("LPOP", db.ListLPOP)
	List.Register("LDEL", db.ListDEL)
	List.Register("LLEN", db.ListLEN)
}

// bigoHashTable namespace
func bigoHashTable() {
	HashTable.Register("HGET", db.HashTableGET)
	HashTable.Register("HSET", db.HashTableSET)
	HashTable.Register("HGETFIELDS", db.HashTableGETFIELDS)
	HashTable.Register("HSETFIELD", db.HashTableSETFIELD)
	HashTable.Register("HSETMULTIFIELDS", db.HashTableSETMULTIFIELDS)
	HashTable.Register("HLEN", db.HashTableLEN)
}

// bigoSet namespace
func bigoSet() {
	Set.Register("SGET", db.SetGET)
	Set.Register("SSET", db.SetSET)
	Set.Register("SPUSH", db.SetPUSH)
	Set.Register("SDEL", db.SetDEL)
	Set.Register("SLEN", db.SetLEN)
}
