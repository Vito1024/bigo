package datastructure

/*
	Doubly linked list implementation
*/

type node struct {
	next  *node
	pre   *node
	value interface{}
}

type List struct {
	head   *node
	tail   *node
	length int
}

func (l *List) Append(value interface{}) {
	node := &node{
		value: value,
	}

	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		node.pre = l.tail
		l.tail = node
	}
	l.length++
}

func (l *List) Delete(value interface{}) int {
	elementDeleted := 0
	for nd := l.head; nd != nil; nd = nd.next {
		if nd.value == value {
			if nd == l.head {
				// head
				l.head = l.head.next
				l.head.pre = nil
			} else if nd == l.tail {
				// tail
				l.tail = l.tail.pre
				l.tail.next = nil
			} else {
				// middle
				nd.pre.next = nd.next
				nd.next.pre = nd.pre
			}
			l.length--
			elementDeleted++
		}
	}
	return elementDeleted
}

func (l *List) GetAll() []interface{} {
	elements := make([]interface{}, 0, l.length)
	for nd := l.head; nd != nil; nd = nd.next {
		elements = append(elements, nd.value)
	}
	return elements
}

func (l *List) Len() int {
	return l.length
}

func (l *List) Pop() interface{} {
	if l.tail == nil {
		return nil
	}
	pop := l.tail
	l.tail = l.tail.pre
	if l.tail == nil {
		// now, it is empty
		l.head = l.tail
	} else {
		l.tail.next = nil
	}
	return pop.value
}

func (l *List) LPop() interface{} {
	if l.head == nil {
		return nil
	}
	pop := l.head
	l.head = l.head.next
	if l.head == nil {
		// now, it's empty
		l.tail = l.head
	} else {
		l.head.pre = nil
	}
	return pop
}

