package model

// If an error happened, return []byte is nil
type Handler func(args []string) ([]byte, error)

// abstract of types supported
type BigoType interface {
	// fetch the function registered from BigoType, if not exists, bool is false
	Fetch(commandName string) (Handler, bool)

	// register a command to current type
	Register(commandName string, handler Handler)
}
