package utils

import (
	"github.com/satori/go.uuid"
)

func Uuid() string {
	u := uuid.NewV4()
	return u.String()
}