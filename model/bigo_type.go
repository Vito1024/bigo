package model

/* 
	The Handler abstracts commands implementation of bigo.


	If an error happened, return []byte is nil
*/
type Handler func(args []string) ([]byte, error)

/* Abstract of bigo types */
type BigoType interface {
	/* 
		Fetch the function registered in it from specific BigoType, 
	if not exists, bool is false 
	*/
	Fetch(commandName string) (Handler, bool)

	/* Register a command to current type */
	Register(commandName string, handler Handler)
}

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
