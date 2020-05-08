package test

import (
	"bigo/utils"
	"fmt"
	"testing"
)

//func TestUtil_SplitString(t *testing.T) {
//	log.Println(strings.Join(utils.SplitString(" hello world hello world hello world  ", ' '), ","))
//}

func TestUtil_Split(t *testing.T) {
	res, err := utils.Split("hello world \"hello world\"",' ')
	if err != nil {
		t.Fail()
	}
	fmt.Println(res, len(res))
}