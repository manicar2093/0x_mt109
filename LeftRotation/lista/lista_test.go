package lista

import (
	"testing"
)

const (
	HeadError = "Head diferente al requerido"
	TailError = "Tail diferente al requerido"
	NextError = "Next diferente al requerido"
)

func createLinkedList() LinkedList {
	lista := LinkedList{}
	lista.Insert(1)
	lista.Insert(2)
	lista.Insert(3)
	lista.Insert(4)
	lista.Insert(5)
	return lista
}

// TestInsertFirst valida que se agregen correctamente los nodos a la lista y se vinculen correctamente
func TestShift(t *testing.T) {
	testTable := []struct {
		Lista      LinkedList
		ShiftTimes int
		Expected   string
	}{
		{Lista: createLinkedList(), ShiftTimes: 0, Expected: "[1,2,3,4,5]"},
		{Lista: createLinkedList(), ShiftTimes: 1, Expected: "[2,3,4,5,1]"},
		{Lista: createLinkedList(), ShiftTimes: 2, Expected: "[3,4,5,1,2]"},
		{Lista: createLinkedList(), ShiftTimes: 3, Expected: "[4,5,1,2,3]"},
		{Lista: createLinkedList(), ShiftTimes: 4, Expected: "[5,1,2,3,4]"},
		{Lista: createLinkedList(), ShiftTimes: 5, Expected: "[1,2,3,4,5]"},
	}

	for i, v := range testTable {
		v.Lista.Shift(v.ShiftTimes)
		got := v.Lista.Display()
		if got != v.Expected {
			t.Errorf("Prueba fallida en indice %d", i)
		}
	}
}
