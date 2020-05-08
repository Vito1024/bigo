package db

import (
	"strconv"

	"bigo/datastructure"
	"bigo/model"
)

func ListGET(args []string) ([]byte, error) {
	if len(args) != 1 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := args[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(*datastructure.List)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}
		_data := data.GetAll()
		res := make([]byte, 0, len(_data))
		for _, v := range _data {
			s := v.(string)
			res = append(res, (s + " ")...)
		}

		return res, nil
	}

	return keyNotFoundMessage, keyNotFoundErr
}

func ListSET(args []string) ([]byte, error) {
	if len(args) < 2 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := args[0]
	if v, ok := BigoDB[key]; ok && v.Type != model.BigoList {
		return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
	}

	data := &datastructure.List{}
	for _, v := range args[1:] {
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

func ListAPPEND(args []string) ([]byte, error) {
	if len(args) < 2 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := args[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(*datastructure.List)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}
		for _, v := range args[1:] {
			data.Append(v)
		}

		return okMessage, nil
	}

	return keyNotFoundMessage, keyNotFoundErr
}

func ListLAPPEND(args []string) ([]byte, error) {
	if len(args) < 2 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := args[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(*datastructure.List)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}
		for _, v := range args[1:] {
			data.LAppend(v)
		}

		return okMessage, nil
	}

	return keyNotFoundMessage, keyNotFoundErr
}

func ListPOP(args []string) ([]byte, error) {
	if len(args) != 1 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := args[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(*datastructure.List)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}
		value := data.Pop()
		if value != nil {
			return []byte(value.(string)), nil
		} else {
			return emptyKeyMessage, emptyKeyErr
		}
	}

	return keyNotFoundMessage, keyNotFoundErr
}

func ListLPOP(args []string) ([]byte, error) {
	if len(args) != 1 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := args[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(*datastructure.List)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}
		value := data.LPop()
		if value != nil {
			return []byte(value.(string)), nil
		} else {
			return emptyKeyMessage, emptyKeyErr
		}
	}

	return keyNotFoundMessage, keyNotFoundErr
}

func ListDEL(args []string) ([]byte, error) {
	if len(args) < 2 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := args[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(*datastructure.List)
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

func ListLEN(args []string) ([]byte, error) {
	if len(args) != 1 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := args[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(*datastructure.List)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}

		return []byte(strconv.Itoa(data.Len())), nil
	}

	return keyNotFoundMessage, keyNotFoundErr
}
