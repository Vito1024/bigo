package db

import (
	"bigo/model"
)

var HashTable = newHashTable()

func newHashTable() *model.HashTable {
	hashtable := &model.HashTable{}

	return hashtable
}


//func HashTableGET(ctx context.Context) {
//
//}
//
//func HashTableSET(ctx context.Context) {
//
//}