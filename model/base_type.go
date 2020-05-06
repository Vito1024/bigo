package model

type baseType struct {
	Commands map[string]Handler
}

func newbaseType() *baseType {
	bt := &baseType{
		Commands: make(map[string]Handler),
	}

	return bt
}

func (bt *baseType) Register(commandName string, handler Handler) {
	bt.Commands[commandName] = handler
}

func (bt *baseType) Fetch(commandName string) (Handler, bool) {
	if handler, ok := bt.Commands[commandName]; ok {
		return handler, true
	}
	return nil, false
}
