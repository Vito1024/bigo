package db

import (
	"bigo/model"
	"encoding/json"
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

