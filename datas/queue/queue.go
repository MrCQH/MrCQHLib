package queue

// TODO 循环队列待实现
// 普通队列 利用append实现
type Queue[T any] struct {
	q []T // 集合
	//hh, tt int   // hh 队头: 出队; tt 队尾: 入队
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		q: make([]T, 0),
		//hh: 0,
		//tt: 0,
	}
}

// 入队
func (o *Queue[T]) Push(v T) {
	o.q = append(o.q, v)
}

// 出队, 如果为空，则为nil
func (o *Queue[T]) Pop() (e T) {
	if o.IsEmpty() {
		return
	}
	e = o.q[0]
	o.q = o.q[1:]
	return
}

// 判空
func (o *Queue[T]) IsEmpty() bool {
	return len(o.q) == 0
}

// 获得队头, 如果为空，则为nil
func (o *Queue[T]) Back() (e T) {
	if o.IsEmpty() {
		return
	}
	e = o.q[0]
	return
}

// 获得队尾, 如果为空，则为nil
func (o *Queue[T]) Front() (e any) {
	if o.IsEmpty() {
		return
	}
	e = o.q[len(o.q)-1]
	return
}

// 队列元素数量
func (o *Queue[T]) Len() int {
	return len(o.q)
}
