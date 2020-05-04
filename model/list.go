package model

type List struct {
	Commands map[string]Handler
	Datas    map[BigoObject]BigoObject
}

func (l *List) Register(commandName string, handler Handler) {
	l.Commands[commandName] = handler
}
