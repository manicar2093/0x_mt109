package main

import (
	"fmt"

	"github.com/manicar2093/leftRotation/lista"
)

func main() {
	l := lista.LinkedList{}

	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(4)
	l.Insert(5)

	fmt.Println(l.Display())

	l.Delete()
	fmt.Println(l.Display())

	l.Shift(2)
	fmt.Println(l.Display())

}
