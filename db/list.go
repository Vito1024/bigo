package db

import (
	"bigo/datastructure"
	"bigo/model"
	"strings"
)

func ListGET(args []byte) ([]byte, error) {
	argStr := strings.Split(string(args), " ")
	if len(argStr) != 1 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := argStr[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(*datastructure.List)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}
		_data := data.GetAll()
		res := make([]byte, 0, len(_data))
		for _, v := range _data {
			s := v.(string)
			res = append(res, (s+" ")...)
		}

		return res, nil
	}

	return keyNotFoundMessage, keyNotFoundErr
}

func ListSET(args []byte) ([]byte, error) {
	argStrs := strings.Split(string(args), " ")
	if len(argStrs) < 2 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := argStrs[0]
	if v, ok := BigoDB[key]; ok && v.Type != model.BigoList {
		return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
	}

	data := &datastructure.List{}
	for _, v := range argStrs[1:] {
		data.Append(v)
	}

	bigoValue := &model.BigoValue{
		Type:     model.BigoList,
		Encoding: model.BigoEncodingList,
		Data:     data,
	}
	BigoDB[key] = bigoValue
	return okMessage, nil
}

func ListAPPEND(args []byte) ([]byte, error) {
	argStr := strings.Split(string(args), " ")
	if len(argStr) < 2 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := argStr[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(*datastructure.List)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}
		for _, v := range argStr[1:] {
			data.Append(v)
		}

		return okMessage, nil
	}

	return keyNotFoundMessage, keyNotFoundErr
}
