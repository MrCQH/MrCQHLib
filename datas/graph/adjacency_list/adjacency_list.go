package ajl

import "MrCQHLib/datas/queue"

// TODO 实现无向图
// 邻接表
type AJL[W any] struct {
	h            map[int]*node[W] // 表头id -> 下一个node
	h2n          map[int]*node[W] // 表头id -> 对应表头node
	r            map[int]struct{} // 存在id，使其不重复搜索点。支持遍历环。每次搜索，记得清空
	cacheTopSort []int            // 缓存上次的topSort
}

type node[W any] struct {
	id int
	w  W        // prev -> cur 边的权重，初始为nil
	ne *node[W] // 下一个点
	d  int      //入度
}

// 初始化
func New[W any]() *AJL[W] {
	return &AJL[W]{
		h:   make(map[int]*node[W]),
		h2n: make(map[int]*node[W]),
		r:   make(map[int]struct{}),
	}
}

// 将u,v建边, 长度为w
func (o *AJL[W]) Add(u, v int, w W) {
	und := &node[W]{id: u}
	vnd := &node[W]{id: v, w: w}
	if _, ok := o.h2n[u]; !ok {
		// 额外记录id -> 表头node
		o.h2n[u] = und
	}
	if _, ok := o.h2n[v]; !ok {
		// 额外记录id -> 表头node
		o.h2n[v] = vnd
	}

	if _, ok := o.h[u]; !ok {
		o.h[u] = vnd
		vnd.d++
	} else {
		// 排除重复的点v
		if o.Contain(u, v) {
			return
		}
		vnd.ne = o.h[u]
		o.h[u].d++
		// 头插
		o.h[u] = vnd
		vnd.d++
	}
}

// 将u->v单向边删除
// 返回 true/false, u, v
// 若为false 两个点返回-1, -1
func (o *AJL[w]) Remove(u, v int) bool {
	if _, ok := o.h[u]; !ok {
		return false
	}
	vnd := o.h[u]
	if vnd.id == v {
		o.h[u] = vnd.ne
		o.h[vnd.id] = nil
		return true
	}
	for e := vnd; e != nil; e = e.ne {
		if e.ne != nil && e.ne.id == v {
			o.h[e.ne.id] = nil
			e.ne = e.ne.ne
			return true
		}
	}
	return false
}

// a节点是否有指向b节点的边
func (o *AJL[W]) Contain(a, b int) bool {
	c, ok := o.h[a]
	if !ok {
		return false
	}
	for ; c != nil; c = c.ne {
		if c.id == b {
			return true
		}
	}
	return false
}

// 判断有无环
func (o *AJL[W]) haveRing() bool {
	limit := len(o.h2n) * 2
	var cnt int
	for _, snd := range o.h {
		for e := snd; e != nil && o.h[e.id] != nil; e = o.h[e.id] {
			cnt++
			if cnt >= limit {
				return true
			}
		}
	}
	return false
}

// 只能遍历一次
func (o *AJL[W]) TopSort() (topList []int) {
	if o.haveRing() {
		return nil
	}
	topList = make([]int, 0)
	q := queue.New[*node[W]]()
	for _, hnd := range o.h2n {
		if hnd.d == 0 {
			q.Push(hnd)
			topList = append(topList, hnd.id)
		}
	}

	for !q.IsEmpty() {
		t := q.Pop()
		for e := o.h[t.id]; e != nil; e = e.ne {
			e.d--
			if e.d == 0 {
				q.Push(e)
				topList = append(topList, e.id)
			}
		}
		//if t.d == 0 && t.ne != nil {
		//	if !f(t.id, t.ne.id, t.ne.w) {
		//		return
		//	}
		//	t.ne.d--
		//	q.Push(t.ne)
		//}
	}
	o.cacheTopSort = make([]int, len(topList))
	copy(o.cacheTopSort, topList)
	return
}

func (o *AJL[W]) DfsRange(u int, f func(a, b int, w W) bool) {
	o.r[u] = struct{}{}
	o.doDfsRange(u, f)
	o.resetR()
}

// 以u为起点，dfs遍历临接表,u->v
func (o *AJL[W]) doDfsRange(u int, f func(a, b int, w W) bool) {
	vnd, ok := o.h[u]
	// 表示该点是最后一个点，就没有必要遍历了
	if !ok || vnd == nil {
		return
	}
	if _, ok = o.r[vnd.id]; ok {
		return
	}
	o.r[vnd.id] = struct{}{}
	for ; vnd != nil; vnd = vnd.ne {
		if !f(u, vnd.id, vnd.w) {
			return
		}
		o.doDfsRange(vnd.id, f)
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
