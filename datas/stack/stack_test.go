package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stk := New[int](3)
	stk.Push(1)
	stk.Push(2)
	stk.Push(3)
	fmt.Println(stk)
}
