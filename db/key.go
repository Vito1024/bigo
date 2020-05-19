package db

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"errors"

	"bigo/model"
)

func KeyDEL(args []string) ([]byte, error) {
	if len(args) < 1 {
		return nil, argsFormatWrongErr
	}

	for _, str := range args {
		delete(BigoDB, str)
	}

	return okMessage, nil
}

func KeyTYPE(args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, argsFormatWrongErr
	}
	key := args[0]

	if v, ok := BigoDB[key]; ok {
	return []byte(v.Type), nil
	} else {
		return nil, keyNotFoundErr
	}
}

func KeyKEY(args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, argsFormatWrongErr
	}
	args[0] = strings.ToLower(args[0])

	res := make([]byte, 0)
	if args[0] == "*" {
		for k := range BigoDB {
			res = append(res, []byte(k+"\n")...)
		}
	} else {
		for k, v := range BigoDB {
			if v.Type == args[0] {
				res = append(res, []byte(k+"\n")...)
			}
		}
	}
	if len(res) != 0 && res[len(res)-1]=='\n' {
		// trim last '\n'
		res = res[:len(res)-1]
	}

	return res, nil
}

func KeyPING(args []string) ([]byte, error) {
	if len(args) != 0 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}

	return []byte("pong"), nil
}

func KeyCOUNT(args []string) ([]byte, error) {
	if len(args) != 0 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}

	return []byte(strconv.Itoa(len(BigoDB))), nil
}

func KeySELECT(args []string) ([]byte, error) {
	if len(args) != 1 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}
	_namespace, err := strconv.Atoi(args[0])
	if err != nil {
		return argsFormatWrongMessage, argsFormatWrongErr
	}

	if _namespace > 15 || _namespace < 0 {
		return []byte("out of range"), errors.New("out of range")
	}
	BigoDB = BigoDBS[_namespace]
	namespace = _namespace
	return okMessage, nil
}

func KeyDB(args []string) ([]byte, error) {
	if len(args) != 0 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}

	return []byte(strconv.Itoa(namespace)), nil
}


func KeyDUMP(args []string) ([]byte, error) {
	if len(args) != 0  {
		return argsFormatWrongMessage, argsFormatWrongErr
	}

	byteStream, err := json.Marshal(BigoDB)
	if err != nil {
		return []byte(err.Error()), err
	}

	err = writeToDisk(byteStream)
	if err != nil {
		return []byte(err.Error()), err
	}
	return okMessage, nil
}

func writeToDisk(byteStream []byte) error {
	if err := ioutil.WriteFile("/var/tmp/bigoDump.json", byteStream, 0644); err != nil {
		return err
	}
	return nil
}

func KeyRECOVER(args []string) ([]byte, error) {
	if len(args) != 0 {
		return argsFormatWrongMessage, argsFormatWrongErr
	}

	file, err := os.Open("/var/tmp/bigoDump.json")
	if err != nil {
		return []byte(err.Error()), err
	}
	defer file.Close()

	byteStream, err := ioutil.ReadAll(file)
	if err != nil {
		return []byte(err.Error()), err
	}

	newDB := make(map[string]*model.BigoValue)
	err = json.Unmarshal(byteStream, &newDB)
	if err != nil {
		return []byte(err.Error()), nil
	}

	BigoDB = newDB
	return okMessage, nil
}
