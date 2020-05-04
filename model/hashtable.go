package model

type HashTable struct {
	Commands map[string]Handler
	Datas map[string]interface{}
}

func(h *HashTable) Register(commandName string, handler Handler) {
	h.Commands[commandName] = handler
}

