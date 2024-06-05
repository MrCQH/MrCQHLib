package ajl

import "MrCQHLib/datas/queue"

// TODO 实现无向图
// 邻接表
type AJL[W any] struct {
	h map[int]*node[W] // id -> node 表头
	r map[int]struct{} // 存在id，使其不重复搜索点。支持遍历环。每次搜索，记得清空
}

type node[W any] struct {
	id int
	w  W        // head -> cur 边的权重，初始为nil
	ne *node[W] // 下一个点
}

// 初始化
func New[W any]() *AJL[W] {
	return &AJL[W]{
		h: make(map[int]*node[W]),
		r: make(map[int]struct{}),
	}
}

// 将u,v建边, 长度为w
func (o *AJL[W]) Add(u, v int, w W) {
	if _, ok := o.h[u]; !ok {
		nd := &node[W]{id: v, w: w}
		o.h[u] = nd
	} else {
		nd := &node[W]{id: v, w: w}
		// 头插
		nd.ne = o.h[u]
		o.h[u] = nd
	}
}

// 将u->v单向边删除
// 返回 true/false, u, v
// 若为false 两个点返回-1, -1
func (o *AJL[w]) Remove(u, v int) bool {
	if _, ok := o.h[u]; !ok {
		return false
	}
	for e := o.h[u]; e != nil; e = e.ne {
		if e.id == v {
			o.h[u] = e.ne
			return true
		}
		if e.ne != nil && e.ne.id == v {
			e.ne = e.ne.ne
			return true
		}
	}
	return false
}

// a节点是否有指向b节点的边
func (o *AJL[W]) Contain(a, b int) bool {
	for c, ok := o.h[a]; c.ne != nil && ok; c = c.ne {
		if c.id == b {
			return true
		}
	}
	return false
}

func (o *AJL[W]) DfsRange(u int, f func(a, b int, w W) bool) {
	o.doDfsRange(u, f)
	o.resetR()
}

// 以u为起点，dfs遍历临接表,u->v
func (o *AJL[W]) doDfsRange(u int, f func(a, b int, w W) bool) {
	v, ok := o.h[u]
	// 表示该点是最后一个点，就没有必要遍历了
	if !ok {
		return
	}
	if _, ok = o.r[v.id]; ok {
		return
	}
	o.r[v.id] = struct{}{}
	for ; v != nil; v = v.ne {
		if !f(u, v.id, v.w) {
			return
		}
		o.doDfsRange(v.id, f)
	}
}

func (o *AJL[W]) BfsRange(u int, f func(a, b int, w W) bool) {
	o.doBfsRange(u, f)
	o.resetR()
}

// 以u为起点,bfs遍历领接表, u->v
func (o *AJL[W]) doBfsRange(u int, f func(a, b int, w W) bool) {
	q := queue.New[int]()
	// 初始点
	q.Push(u)
	o.r[u] = struct{}{}

	// 层次遍历
	for !q.IsEmpty() {
		t := q.Pop()
		for e := o.h[t]; e != nil; e = e.ne {
			if _, ok := o.r[e.id]; !ok {
				o.r[e.id] = struct{}{}
				if !f(t, e.id, e.w) {
					return
				}
				q.Push(e.id)
			}
		}
	}
}

// 手动清除
func (o *AJL[W]) resetR() {
	for id := range o.r {
		delete(o.r, id)
	}
}
