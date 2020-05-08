package db

import (
	"bigo/datastructure"
	"bigo/model"
	"strconv"
)

func SetGET(args []string) ([]byte, error) {
	if len(args) != 1 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}

	key := args[0]
	if v, ok := BigoDB[key]; ok {
		_data, ok := v.Data.(*datastructure.Set)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}

		return _data.GetAll(), nil
	}

	return keyNotFoundMessage, keyNotFoundErr
}

func SetSET(args []string) ([]byte, error) {
	if len(args) < 2 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}

	data := datastructure.NewSet()
	for _, v := range args[1:] {
		data.Push(v)
	}

	key := args[0]
	if v, ok := BigoDB[key]; ok && v.Type != model.BigoSet {
		return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
	}

	bigoValue := &model.BigoValue{
		Type:     model.BigoSet,
		Encoding: model.BigoEncodingSet,
		Data:     data,
	}

	BigoDB[key] = bigoValue

	return okMessage, nil
}

func SetDEL(args []string) ([]byte, error) {
	if len(args) < 2 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := args[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(*datastructure.Set)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}

		for _, str := range args[1:] {
			data.Delete(str)
		}
		return okMessage, nil
	}

	return keyNotFoundMessage, keyNotFoundErr
}

func SetPUSH(args []string) ([]byte, error) {
	if len(args) < 2 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := args[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(*datastructure.Set)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}

		for _, str := range args[1:] {
			data.Push(str)
		}

		return okMessage, nil
	}

	return keyNotFoundMessage, keyNotFoundErr
}

func SetLEN(args []string) ([]byte, error) {
	if len(args) != 1 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := args[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(*datastructure.Set)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}

		return []byte(strconv.Itoa(data.Len())), nil
	}

	return keyNotFoundMessage, keyNotFoundErr
}
