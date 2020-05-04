package datastructure

import (
	"math/rand"
	"testing"
)

func TestList_Append(t *testing.T) {
	l := List{}
	for i := 0; i < 10; i++ {
		l.Append(i)
	}

	for i, v := range l.GetAll() {
		if i != v.(int) {
			t.Fail()
		}
	}
}

func TestList_Delete(t *testing.T) {
	l := List{}
	randNum := 0
	for randNum < 5 {
		// At least 5 elements
		randNum = rand.Intn(30) // [0, 30)
	}
	for i := 0; i < randNum; i++ {
		l.Append(i)
	}
	t.Log(randNum, "elements appended")

	t.Log("**** delete head ****")
	l.Delete(0)
	t.Log(l.GetAll())

	i := 1
	nd := l.head
	for ; i < randNum && nd != nil; i, nd = i+1, nd.next {
		if nd.value.(int) != i {
			t.Fail()
		}
	}
	if i != randNum || nd != nil {
		t.Fail()
	}

	t.Log("**** delete tail ****")
	l.Delete(randNum-1)
	t.Log(l.GetAll())

	i = 1
	nd = l.head
	for ; i < randNum-1 && nd != nil; i, nd = i+1, nd.next {
		if nd.value.(int) != i {
			t.Fail()
		}
	}
    if i != randNum-1 || nd != nil {
    	t.Fail()
	}

	t.Log("**** delete middle ****")
	l.Delete(randNum/2)
	t.Log(randNum/2, "is deleted")
	t.Log(l.GetAll())
}

func TestList_GetAll(t *testing.T) {
	l := List{}
	randNum := rand.Intn(30) // [0, 30)
	for i := 0; i < randNum; i++ {
		l.Append(i)
	}
	t.Log(randNum, "elements appended")
	t.Log(l.GetAll())
}

func TestList_Len(t *testing.T) {
	l := List{}
	randNum := rand.Intn(100)
	for i := 0; i < randNum; i++ {
		l.Append(i)
	}
	t.Log(randNum, "elements appended")
	if l.Len() != randNum {
		t.Fail()
	}
	elementDeleted := l.Delete(0)
	t.Log(elementDeleted, "elements deleted")
	if l.Len() != randNum - elementDeleted {
		t.Fail()
	}
}