package stack

// TODO 差错误检查和测试
type Stack[T any] struct {
	elems []T
	rp    int
}

func NewStack[T any](cap ...int) *Stack[T] {
	var size int
	if len(cap) > 0 {
		size = cap[0]
	}
	return &Stack[T]{
		elems: make([]T, size),
		rp:    -1,
	}
}

// 入栈
func (s *Stack[T]) Push(v T) {
	s.rp += 1
	s.elems[s.rp] = v
}

// 出栈
func (s *Stack[T]) Pop() (e T) {
	e = s.elems[s.rp]
	s.rp -= 1
	return
}

// 栈容量
func (s *Stack[T]) Len() int {
	return len(s.elems)
}

// 栈顶
func (s *Stack[T]) Peek() T {
	return s.elems[s.rp]
}
