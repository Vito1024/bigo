package db

import (
	"bigo/model"
	"encoding/json"
	"strconv"
)

func HashTableGET(args []string) ([]byte, error) {
	if len(args) != 1 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}

	key := args[0]
	if v, ok := BigoDB[key]; ok {
		if v.Type != model.BigoHashTable {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}
		return json.Marshal(v.Data)
	}

	return keyNotFoundMessage, keyNotFoundErr
}

func HashTableGETFIELDS(args []string) ([]byte, error) {
	if len(args) < 2 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}

	mapBytes, err := HashTableGET(args[:1])
	if err != nil {
		return mapBytes, err
	}

	data := make(map[string]string)
	if err := json.Unmarshal(mapBytes, &data); err != nil {
		return []byte(err.Error()), err
	}

	var res = make([]byte, 0, 20)
	for _, k := range args[1:] {
		if v, ok := data[k]; ok {
			res = append(res, (v+"\n")...)
		} else {
			// field not found
			res = append(res, "\n"...)
		}
	}

	if len(res) > 0 && res[len(res)-1] == '\n' {
		res = res[:len(res)-1]
	}
	return res, nil
}

func HashTableSET(args []string) ([]byte, error) {
	if len(args) < 3 || len(args) % 2 == 0 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}

	hashTable := make(map[string]string, len(args)/2)
	for i := 1; i < len(args); i += 2 {
		hashTable[args[i]] = args[i+1]
	}

	key := args[0]
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

func HashTableSETFIELD(args []string) ([]byte, error) {
	if len(args) != 3 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := args[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(map[string]string)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}

		data[args[1]] = data[args[2]]
		return okMessage, nil
	}

	return keyNotFoundMessage, keyNotFoundErr
}

func HashTableSETMULTIFIELDS(args []string) ([]byte, error) {
	if len(args) < 3 && len(args)%2 == 0 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := args[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(map[string]string)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}

		for i := 2; i < len(args); i += 2 {
			data[args[i-1]] = args[i]
		}
		return okMessage, nil
	}

	return keyNotFoundMessage, keyNotFoundErr
}


func HashTableLEN(args []string) ([]byte, error) {
	if len(args) != 1 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := args[0]

	if v, ok := BigoDB[key]; ok {
		data, ok := v.Data.(map[string]string)
		if !ok {
			return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
		}

		return []byte(strconv.Itoa(len(data))), nil
	}

	return keyNotFoundMessage, keyNotFoundErr
}
