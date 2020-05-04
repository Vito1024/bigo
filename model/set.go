package model

type Set struct {
	Commands map[string]Handler
}

func(s *Set) Register(commandName string, handler Handler) {
	s.Commands[commandName] = handler
}

