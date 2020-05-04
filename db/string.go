package db

import (
	"bigo/model"
	"bytes"
	"errors"
	"unsafe"
)

var String = newString()

func newString() *model.String {
	str := &model.String{
		Commands: make(map[string]model.Handler),
		Datas:    make(map[string]model.BigoObject),
	}

	return str
}

func StringGET(args []byte) ([]byte, error) {
	bytesSlice := bytes.Split(args, []byte{' '})
	if len(bytesSlice) != 1 {
		return nil, errors.New("string GET command format error: length != 1")
	}

	getFormat := model.GETFormat{Key: string(args)}

	key := model.BigoObject{
		Type:     model.BigoString,
		Encoding: model.BigoString,
		Ptr:      unsafe.Pointer(&getFormat.Key),
	}
	resObject, err := stringGET(key)
	if err != nil {
		return nil, err
	}

	res := (*string)(resObject.Ptr)
	return []byte(*res+"\n"), nil
}

func stringGET(key model.BigoObject) (model.BigoObject, error) {
	keyString := *(*string)(key.Ptr)
	if v, ok := String.Datas[keyString]; ok {
		return v, nil
	}
	return model.BigoObject{}, errors.New("key not exists")
}

func StringSET(args []byte) ([]byte, error) {
	bytesSlice := bytes.Split(args, []byte{' '})
	if len(bytesSlice) != 2 {
		return nil, errors.New("string SET args format error: length != 2")
	}
	setFormat := model.SETFormat{
		Key: string(bytesSlice[0]),
		Value:string(bytesSlice[1]),
	}

	key := model.BigoObject{
		Type:     model.BigoString,
		Encoding: model.BigoString,
		Ptr:      unsafe.Pointer(&setFormat.Key),
	}

	value := model.BigoObject{
		Type: model.BigoString,
		Encoding:model.BigoString,
		Ptr: unsafe.Pointer(&setFormat.Value),
	}

	stringSET(key, value)
	return []byte("OK\n"), nil
}

func stringSET(key, value model.BigoObject) {
	keyString := *(*string)(key.Ptr)
	String.Datas[keyString] = value
}
