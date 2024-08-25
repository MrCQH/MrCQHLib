package list

import (
	"encoding/json"
	"errors"
)

// 双向链表
type List[T any] struct {
	Head *node[T] `json:"head"` // 头尾结点
	Tail *node[T] `json:"tail"` // 头尾结点
	len  int
}

type node[T any] struct {
	Prev *node[T] `json:"-"`
	Ne   *node[T] `json:"ne"`
	V    T        `json:"v"`
}

func New[T any]() *List[T] {
	head := new(node[T])
	tail := new(node[T])
	head.Ne = tail
	tail.Prev = head
	return &List[T]{
		Head: head,
		Tail: tail,
		len:  0,
	}
}

// 获得第k个点的值
func (o *List[T]) Get(k int) (v T, err error) {
	k = k + 1 // 让k映射为下标为0
	var curNode *node[T]
	var cnt int
	if k/2 <= o.len {
		curNode = o.Head
		cnt = k
		for cnt > 0 {
			if curNode.Ne == o.Tail {
				err = errors.New("list len less than k")
				return
			}
			curNode = curNode.Ne
			cnt--
		}
	} else {
		curNode = o.Tail
		cnt = o.len - k
		for cnt > 0 {
			if curNode.Prev == o.Head {
				err = errors.New("list len less than k")
				return
			}
			curNode = curNode.Prev
			cnt--
		}
	}

	v = curNode.V
	return
}

// 头插
func (o *List[T]) HeadInsert(v T) {
	nd := &node[T]{o.Head, o.Head.Ne, v}
	nd.Prev.Ne = nd
	nd.Ne.Prev = nd
	o.len++
}

// 尾插
func (o *List[T]) TailInsert(v T) {
	nd := &node[T]{o.Tail.Prev, o.Tail, v}
	nd.Prev.Ne = nd
	nd.Ne.Prev = nd
	o.len++
}

// 插入到第K个点后面的点
func (o *List[T]) Insert(k int, v T) error {
	k = k + 1
	nd := &node[T]{nil, nil, v}
	var curNode *node[T]
	var cnt int

	if k/2 <= o.len {
		curNode = o.Head
		cnt = k
		for cnt > 0 {
			if curNode.Ne == o.Tail {
				return errors.New("list len less than k")
			}
			curNode = curNode.Ne
			cnt--
		}
	} else {
		curNode = o.Tail
		cnt = o.len - k
		for cnt > 0 {
			if curNode.Prev == o.Head {
				return errors.New("list len less than k")
			}
			curNode = curNode.Prev
			cnt--
		}

	}
	nd.Prev = curNode
	nd.Ne = curNode.Ne
	curNode.Ne.Prev = nd
	curNode.Ne = nd
	o.len++
	return nil
}

// 删除第k个点
func (o *List[T]) Remove(k int) error {
	k = k + 1
	var curNode *node[T]
	var cnt int

	if k/2 <= o.len {
		curNode = o.Head
		cnt = k
		for cnt-1 > 0 {
			if curNode.Ne == o.Tail {
				return errors.New("list len less than k")
			}
			curNode = curNode.Ne
			cnt--
		}
	} else {
		curNode = o.Tail
		cnt = o.len - k
		for cnt-1 > 0 {
			if curNode.Prev == o.Head {
				return errors.New("list len less than k")
			}
			curNode = curNode.Prev
			cnt--
		}
	}
	curNode.Ne = curNode.Ne.Ne
	curNode.Ne.Prev = curNode
	o.len--
	return nil
}

func (o *List[T]) Len() int {
	return o.len
}

func (o *List[T]) String() string {
	if bytes, err := json.Marshal(o); err == nil {
		return string(bytes)
	} else {
		return err.Error()
	}
}
