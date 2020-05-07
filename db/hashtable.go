package db

import (
	"bigo/model"
	"encoding/json"
	"strconv"
	"strings"
)

func HashTableGET(args []byte) ([]byte, error) {
	argsStrs := strings.Split(string(args), " ")
	if len(argsStrs) != 1 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}

	key := argsStrs[0]
	if v, ok := BigoDB[key]; ok {
		if v.Type != model.BigoHashTable {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}
		return json.Marshal(v.Data)
	}

	return keyNotFoundMessage, keyNotFoundErr
}

func HashTableGETFIELDS(args []byte) ([]byte, error) {
	argsStr := strings.Split(string(args), " ")
	if len(argsStr) < 2 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}

	mapBytes, err := HashTableGET([]byte(argsStr[0]))
	if err != nil {
		return mapBytes, err
	}

	data := make(map[string]string)
	if err := json.Unmarshal(mapBytes, &data); err != nil {
		return []byte(err.Error()), err
	}

	var res = make([]byte, 0, 20)
	for _, k := range argsStr[1:] {
		if v, ok := data[k]; ok {
			res = append(res, (v+"\n")...)
		} else {
			// field not found
			res = append(res, "\n"...)
		}
	}
	return res, nil
}

func HashTableSET(args []byte) ([]byte, error) {
	argsStrs := strings.Split(string(args), " ")
	if len(argsStrs) < 3 || len(argsStrs) % 2 == 0 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}

	hashTable := make(map[string]string, len(argsStrs)/2)
	for i := 1; i < len(argsStrs); i += 2 {
		hashTable[argsStrs[i]] = argsStrs[i+1]
	}

	key := argsStrs[0]
	if v, ok := BigoDB[key]; ok && v.Type != model.BigoHashTable {
		return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
	}
	BigoDB[key] = &model.BigoValue{
		Type: model.BigoHashTable,
		Encoding:model.BigoEncodingHashTable,
		Data:hashTable,
	}

	return okMessage, nil
}

func HashTableSETFIELD(args []byte) ([]byte, error) {
	argStrs := strings.Split(string(args), " ")
	if len(argStrs) != 3 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := argStrs[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(map[string]string)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}

		data[argStrs[1]] = data[argStrs[2]]
		return okMessage, nil
	}

	return keyNotFoundMessage, keyNotFoundErr
}

func HashTableSETMULTIFIELDS(args []byte) ([]byte, error) {
	argStr := strings.Split(string(args), " ")
	if len(argStr) < 3 && len(argStr)%2 == 0 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := argStr[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(map[string]string)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}

		for i := 2; i < len(argStr); i += 2 {
			data[argStr[i-1]] = argStr[i]
		}
		return okMessage, nil
	}

	return keyNotFoundMessage, keyNotFoundErr
}


func HashTableLEN(args []byte) ([]byte, error) {
	argStr := strings.Split(string(args), " ")
	if len(argStr) != 1 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := argStr[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(map[string]string)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}

		return []byte(strconv.Itoa(len(data))), nil
	}

	return keyNotFoundMessage, keyNotFoundErr
}
