package datastructure

import (
	"math/rand"
	"testing"
)

func TestSet_Append(t *testing.T) {
	s := NewSet()
	randNum := rand.Intn(30)
	for i := 0; i < randNum; i++ {
		s.Append(i)
	}
	t.Log(randNum, "elements appended")
	t.Log(s.GetAll())
}

func TestSet_Delete(t *testing.T) {
	s := NewSet()
	randNum := rand.Intn(30)
	for i := 0; i < randNum; i++ {
		s.Append(i)
	}
	t.Log(randNum, "elements appended")
	s.Delete(randNum/2)
	t.Log(randNum/2, "is deleted")
	if s.In(randNum/2) {
		t.Fail()
	}
	t.Log(s.GetAll())
}

func TestSet_GetAll(t *testing.T) {
	s := NewSet()
	randNum := rand.Intn(30)
	for i := 0; i < randNum; i++ {
		s.Append(i)
	}
	t.Log(randNum, "elements appended")

	t.Log(s.GetAll())
}

func TestSet_Len(t *testing.T) {
	s := NewSet()
	randNum := rand.Intn(30)
	for i := 0; i < randNum; i++ {
		s.Append(i)
	}
	t.Log(randNum, "elements appended")

	if s.Len() != uint(randNum) {
		t.Fail()
	}
}