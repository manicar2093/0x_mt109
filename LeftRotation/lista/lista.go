package lista

import (
	"fmt"
	"strings"
)

type node struct {
	next *node
	data int
}

type LinkedList struct {
	head   *node
	length int
}

func (l *LinkedList) Insert(data int) {
	n := node{data: data}
	var current = &node{}

	if l.head == nil {
		l.head = &n
	} else {
		current = l.head

		for current.next != nil {
			current = current.next
		}

		current.next = &n

	}

	l.length++
}

func (l *LinkedList) Delete() {
	if l.length < 1 {
		return
	}
	previous := &node{}
	current := l.head
	for current.next != nil {
		previous = current
		current = current.next
	}
	previous.next = nil
	l.length--
	return
}

func (l *LinkedList) Display() string {
	var builder strings.Builder

	current := l.head
	count := 0
	builder.WriteString("[")
	for count < l.length {
		builder.WriteString(fmt.Sprintf("%d", current.data))
		count++
		if count != l.length {
			builder.WriteString(",")
		}
		current = current.next
	}
	builder.WriteString("]")

	return builder.String()

}

func (l *LinkedList) Shift(times int) {
	for i := 0; i < times; i++ {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = &node{data: l.head.data}
		l.head = l.head.next
	}
}
