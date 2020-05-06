package db

import (
	"bigo/datastructure"
	"bigo/model"
	"strings"
)

func SetGET(args []byte) ([]byte, error) {
	argStrs := strings.Split(string(args), " ")
	if len(argStrs) != 1 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}

	key := argStrs[0]
	if v, ok := BigoDB[key]; ok {
		_data, ok := v.Data.(*datastructure.Set)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}

		return _data.GetAll(), nil
	}

	return keyNotFoundMessage, keyNotFoundErr
}

func SetSET(args []byte) ([]byte, error) {
	argStrs := strings.Split(string(args), " ")
	if len(argStrs) < 2 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}

	data := datastructure.NewSet()
	for _, v := range argStrs[1:] {
		data.Append(v)
	}

	key := argStrs[0]
	if v, ok := BigoDB[key]; ok && v.Type != model.BigoSet {
		return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
	}

	bigoValue := &model.BigoValue{
		Type: model.BigoSet,
		Encoding: model.BigoEncodingSet,
		Data: data,
	}

	BigoDB[key] = bigoValue

	return okMessage, nil
}