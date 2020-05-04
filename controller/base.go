package controller

import (
	"bigo/db"
	"bigo/model"
)

/*
	Abstract data type supported to BigoType interface
*/
var Bigo map[uint8]model.BigoType // Top level namespace

func init() {

	Bigo = make(map[uint8]model.BigoType)

	Bigo[model.BigoString] = db.String
	//Bigo[model.BigoList] = db.List
	//Bigo[model.BigoHashTable] = db.HashTable
	//Bigo[model.BigoSet] = db.Set

	// bigoString
	bigoString()

	// bigoList
	//bigoList()
	//
	//// bigoHashTable
	//bigoHashTable()
	//
	//// bigoSet
	//bigoSet()
}

// bigoString namespace
func bigoString() {
	db.String.Register("GET", db.StringGET)
	db.String.Register("SET", db.StringSET)
}

// bigoList namespace
//func bigoList() {
//	db.List.Register("LGET", db.ListGET)
//	db.List.Register("LSET", db.ListSET)
//	db.List.Register("LAPPEND", db.ListAPPEND)
//}
//
//// bigoHashTable namespace
//func bigoHashTable() {
//	db.HashTable.Register("HGET", db.HashTableGET)
//	db.HashTable.Register("HSET", db.HashTableSET)
//}
//
//// bigoSet namespace
//func bigoSet() {
//	db.Set.Register("SGET", db.SetGET)
//	db.Set.Register("SSET", db.SetSET)
//}
