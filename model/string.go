package model

import "log"

type String struct {
	Commands map[string]Handler
	Datas    map[string]BigoObject
}

func (s *String) Register(commandName string, handler Handler) {
	s.Commands[commandName] = handler
}

func (s *String) Call(commandName string, args []byte) []byte {
	function, ok := s.Commands[commandName]
	if !ok {
		log.Fatal("Command not supported")
	}
	res, err := function(args)
	if err != nil {
		panic(err)
	}
	return res
}

// format of string GET command
type GETFormat struct {
	Key string `json:"key"`
}

// format of string SET command
type SETFormat struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
