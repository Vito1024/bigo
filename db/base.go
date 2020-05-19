package db

import (
	"errors"

	"bigo/model"
)

/*
	In bigo, both key and value are model.BigoObject type.
	In db package, all operations are operated on model.BigoObject
*/

const (
	dbCount = 16
)

// Global BigoDB object represents the top level namespace, current database
var (
	BigoDBS [dbCount]map[string]*model.BigoValue
	BigoDB map[string]*model.BigoValue
	namespace = 0
)

func init() {
	for i := 0; i < dbCount; i++ {
		BigoDBS[i] = make(map[string]*model.BigoValue)
	}

	BigoDB = BigoDBS[namespace]
}

var (
	// Response message
	okMessage                              = []byte("OK")
	keyNotFoundMessage                     = []byte("key not found")
	keyTypeErrMessage                      = []byte("key type is wrong")
	argsFormatWrongMessage                 = []byte("args format is wrong")
	keyAlreadyExistsButTypeNotMatchMessage = []byte("key already exists but type not match")
	emptyKeyMessage                        = []byte("empty key")

	// Error
	keyNotFoundErr                    = errors.New(string(keyNotFoundMessage))
	keyTypeErr                        = errors.New(string(keyTypeErrMessage))
	argsFormatWrongErr                = errors.New(string(argsFormatWrongMessage))
	keyAlreadyExistButTypeNotMatchErr = errors.New(string(keyAlreadyExistsButTypeNotMatchMessage))
	emptyKeyErr                       = errors.New(string(emptyKeyMessage))
)
