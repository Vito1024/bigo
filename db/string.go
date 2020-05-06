package db

import (
	"bigo/model"
	"bytes"
	"strings"
)


func StringGET(args []byte) ([]byte, error) {
	bytesSlice := bytes.Split(args, []byte{' '})
	if len(bytesSlice) != 1 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}

	bigoValue, ok := BigoDB[string(args)]
	if !ok {
		return keyNotFoundMessage, keyNotFoundErr
	}

	data, ok := bigoValue.Data.(string)
	if !ok {
		return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
	}

	return []byte(data), nil
}

func StringSET(args []byte) ([]byte, error) {
	strs := strings.Split(string(args), " ")
	if len(strs) != 2 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}

	key := strs[0]
	if v, ok := BigoDB[key]; ok && v.Type != model.BigoString {
		return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
	}

	value := &model.BigoValue{
		Type:     model.BigoString,
		Encoding: model.BigoEncodingString,
		Data:     strs[1],
	}

	BigoDB[key] = value

	return okMessage, nil
}
