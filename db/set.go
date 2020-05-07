package db

import (
	"bigo/datastructure"
	"bigo/model"
	"strconv"
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
		data.Push(v)
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

func SetDEL(args []byte) ([]byte, error) {
	argStrs := strings.Split(string(args), " ")
	if len(argStrs) < 2 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := argStrs[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(*datastructure.Set)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}

		for _, str := range argStrs[1:] {
			data.Delete(str)
		}
		return okMessage, nil
	}

	return keyNotFoundMessage, keyNotFoundErr
}

func SetPUSH (args []byte) ([]byte, error) {
	argStrs := strings.Split(string(args), " ")
	if len(argStrs) < 2 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := argStrs[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(*datastructure.Set)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}

		for _, str := range argStrs[1:] {
			data.Push(str)
		}

		return okMessage, nil
	}

	return keyNotFoundMessage, keyNotFoundErr
}

func SetLEN(args []byte) ([]byte, error) {
	argStrs := strings.Split(string(args), " ")
	if len(argStrs) != 1 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := argStrs[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(*datastructure.Set)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}

		return []byte(strconv.Itoa(data.Len())), nil
	}

	return keyNotFoundMessage, keyNotFoundErr
}