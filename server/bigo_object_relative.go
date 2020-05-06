package server

import (
	"bigo/controller"
	"bigo/model"
	"errors"
)

func fetchHandler(cmdName string) (model.Handler, error) {
	for _, v := range controller.Bigo {
		if handler, ok := v.Fetch(cmdName); ok {
			return handler, nil
		}

	}
	return nil, errors.New("command not supported")
}
