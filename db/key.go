package db

import (
	"strings"
)

func KeyDEL(args []byte) ([]byte, error) {
	strs := strings.Split(string(args), " ")
	if len(strs) < 1 {
		return nil, argsFormatWrongErr
	}

	for _, str := range strs {
		delete(BigoDB, str)
	}

	return okMessage, nil
}

func KeyTYPE(args []byte) ([]byte, error) {
	strs := strings.Split(string(args), " ")
	if len(strs) != 1 {
		return nil, argsFormatWrongErr
	}
	key := strs[0]

	if v, ok := BigoDB[key]; ok {
		return []byte(v.Type), nil
	} else {
		return nil, keyNotFoundErr
	}
}

func KeyKEY(args []byte) ([]byte, error) {
	strs := strings.Split(string(args), " ")
	if len(strs) != 1 {
		return nil, argsFormatWrongErr
	}
	strs[0] = strings.ToLower(strs[0])

	res := make([]byte, 0)
	if strs[0] == "*" {
		for k := range BigoDB {
			res = append(res, []byte(k+"\n")...)
		}
	} else {
		for k, v := range BigoDB {
			if v.Type == strs[0] {
				res = append(res, []byte(k+"\n")...)
			}
		}
	}

	return res, nil
}
