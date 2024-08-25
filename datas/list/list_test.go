package list

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	list := New[int]()
	list.TailInsert(1)
	list.HeadInsert(0)
	list.TailInsert(2)
	list.TailInsert(3)
	fmt.Println(list.Get(0))

	list.Insert(2, 4)

	for p := list.Head.Ne; p != list.Tail; p = p.Ne {
		fmt.Print(p.V)
	}
	fmt.Println()
	fmt.Println(list.Len())

	list.Remove(0)
	for p := list.Head.Ne; p != list.Tail; p = p.Ne {
		fmt.Print(p.V)
	}
	fmt.Println()
	fmt.Println(list.Len())

	fmt.Println(list)
}
