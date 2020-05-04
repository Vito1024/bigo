package model

type Handler func(args []byte) ([]byte, error)

// abstract of types supported
type BigoType interface {
	// Call the function registered in current type
	Call(commandName string, args []byte) []byte

	// register a command to current type
	Register(commandName string, handler Handler)
}
