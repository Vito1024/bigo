package server

import (
	"bigo/model"
)

func typeByBigoRequest(req model.BigoRequest) (uint8, error) {
	return model.BigoString, nil
}
