package db

import (
	"bigo/model"
)


func StringGET(args []string) ([]byte, error) {
	if len(args) != 1 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}

	bigoValue, ok := BigoDB[args[0]]
	if !ok {
		return keyNotFoundMessage, keyNotFoundErr
	}

	data, ok := bigoValue.Data.(string)
	if !ok {
		return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
	}

	return []byte(data), nil
}

func StringSET(args []string) ([]byte, error) {
	if len(args) != 2 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}

	key := args[0]
	if v, ok := BigoDB[key]; ok && v.Type != model.BigoString {
		return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
	}

	value := &model.BigoValue{
		Type:     model.BigoString,
		Encoding: model.BigoEncodingString,
		Data:     args[1],
	}

	BigoDB[key] = value

	return okMessage, nil
}
