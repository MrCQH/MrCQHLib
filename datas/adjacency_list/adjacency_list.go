package ajl

// 邻接表
type AJL struct {
	ns         []*node
	isDirected bool // 有向图还是无向图
}

type node struct {
	id    any
	neLen int // cur -> next 边的权重
	ne    *node
}

// 初始化
func (o *AJL) Init() {

}

// 将u,v建边, 长度为w
func (o *AJL) Add(u any, v any, w int) {

}

// 将u,v边删除
func (o *AJL) Remove(u any, v any) {

}

// 遍历领接表
func (o *AJL) Range(f func(any) bool) {

}
