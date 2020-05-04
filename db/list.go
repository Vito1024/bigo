package db

import (
	"bigo/model"
)

var List = newList()

func newList() *model.List {
	list := &model.List{
		Commands: make(map[string]model.Handler),
		Datas:    make(map[model.BigoObject]model.BigoObject),
	}

	return list
}

//func ListGET(key model.BigoObject) {
//
//}
//
//func ListSET(ctx context.Context) {
//
//}
//
//func ListAPPEND(key, value model.BigoObject) {
//	list := (*datastructure.List)(List.Datas[key].Ptr)
//	list.Append(*value.Ptr)
//}