package utils

import (
	"crypto/md5"
)

type HashValue struct {

}

func Hash(str string) [16]byte {
	return md5.Sum([]byte(str))
}
