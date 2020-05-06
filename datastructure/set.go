package datastructure

/*
	Set implementation in Go
*/

type Set struct {
	values map[interface{}]bool
	length uint
}

func NewSet() *Set {
	return &Set{
		values: make(map[interface{}]bool),
		length: 0,
	}
}

func (s *Set) Append(v interface{}) {
	if _, ok := s.values[v]; ok {
		// already exists, return
		return
	}
	s.values[v] = true
	s.length++
}

func (s *Set) Delete(v interface{}) {
	if _, ok := s.values[v]; !ok {
		// not exists, return
		return
	}
	delete(s.values, v)
	s.length--
}

func (s *Set) Len() uint {
	return s.length
}

func (s *Set) GetAll() []byte {
	elements := make([]byte, 0, s.length)
	for k := range s.values {
		elements = append(elements, (k.(string)+" ")...)
	}
	return elements
}

func (s *Set) In(value interface{}) bool {
	if _, ok := s.values[value]; ok {
		return true
	}
	return false
}
