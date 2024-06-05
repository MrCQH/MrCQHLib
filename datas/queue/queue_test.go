package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := New[int]()
	fmt.Println(q.IsEmpty())
	q.Push(1)
	q.Push(2)
	q.Push(23)
	fmt.Println(q)
	e := q.Pop()
	fmt.Println(e, q)
	fmt.Println(q.IsEmpty())

	fmt.Println(q)
	fmt.Println(q.Front(), q.Back())
	fmt.Println(q)

	for !q.IsEmpty() {
		q.Pop()
	}
	fmt.Println(q.Pop())
}
