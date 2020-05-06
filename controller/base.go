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

	String = model.NewString()
	HashTable = model.NewHashTable()
)

func init() {

	Bigo[model.BigoKey] = Key
	//
	Bigo[model.BigoString] = String
	//BigoDB[model.BigoList] = db.List
	Bigo[model.BigoHashTable] = HashTable
	//BigoDB[model.BigoSet] = db.Set

	// bigoKey
	bigoKey()

	// bigoString
	bigoString()

	// bigoList
	//bigoList()
	//
	// bigoHashTable
	bigoHashTable()
	//
	//// bigoSet
	//bigoSet()
}

// bigoKey namespace
func bigoKey() {
	Key.Register("DEL", db.KeyDEL)
	Key.Register("TYPE", db.KeyTYPE)
	Key.Register("KEY", db.KeyKEY)
}

// bigoString namespace
func bigoString() {
	String.Register("GET", db.StringGET)
	String.Register("SET", db.StringSET)
}

// bigoList namespace
//func bigoList() {
//	db.List.Register("LGET", db.ListGET)
//	db.List.Register("LSET", db.ListSET)
//	db.List.Register("LAPPEND", db.ListAPPEND)
//}
//
// bigoHashTable namespace
func bigoHashTable() {
	HashTable.Register("HGET", db.HashTableGET)
	HashTable.Register("HSET", db.HashTableSET)
}
//
//// bigoSet namespace
//func bigoSet() {
//	db.Set.Register("SGET", db.SetGET)
//	db.Set.Register("SSET", db.SetSET)
//}
