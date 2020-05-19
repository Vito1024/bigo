package db

import (
	"bigo/datastructure"
	"bigo/model"
	"strconv"
	"strings"
)

func BTGET(args []string) ([]byte, error) {
	if len(args) != 1 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	if v, ok := BigoDB[args[0]]; ok {
		if data, ok := v.Data.(*datastructure.BinaryTree); ok {
			ints := data.GetAll()
			strs := make([]string, 0, len(ints))
			for _, v := range ints {
				strs = append(strs, strconv.Itoa(v))
			}
			return []byte(strings.Join(strs, " ")), nil
		}
		return keyAlreadyExistsButTypeNotMatchMessage, keyAlreadyExistButTypeNotMatchErr
	}
	return keyNotFoundMessage, keyNotFoundErr
}

func BTSET(args []string) ([]byte, error) {
	if len(args) < 1 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	key := args[0]
	numbers := make([]int, 0, len(args)-1)
	for _, v := range args[1:] {
		intofv, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, intofv)
	}
	bt := datastructure.NewBinaryTree()
	bt.Set(numbers...)

	bigoValue := &model.BigoValue{
		Type:"BinaryTree",
		Encoding:0,
		Data: bt,
	}
	BigoDB[key] = bigoValue
	return okMessage, nil
}