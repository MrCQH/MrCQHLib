package skip_list

import (
	"math/rand"
)

const (
	maxLevel = 8   // 跳表最大高度
	rate     = 0.5 // 创建概率
)

type SkipList struct {
	head *Node // 头结点
	lev  int
}

type Node struct {
	ne []*Node
	v  int
}

func New() SkipList {
	head := new(Node) // 增加头结点
	head.ne = make([]*Node, maxLevel)
	return SkipList{
		head: head,
		lev:  maxLevel,
	}
}

func (this *SkipList) Search(target int) bool {
	prevNodes := this.find(target)
	return prevNodes[0].ne[0] != nil && prevNodes[0].ne[0].v == target
}

func (this *SkipList) Add(target int) {
	prevNodes := this.find(target)
	p := &Node{
		ne: make([]*Node, this.lev),
		v:  target,
	}
	for i := 0; i < this.lev; i++ {
		p.ne[i] = prevNodes[i].ne[i]
		prevNodes[i].ne[i] = p
		if rand.Float64() > rate {
			break
		}
	}
}

func (this *SkipList) Erase(target int) bool {
	prevNodes := this.find(target)
	if prevNodes[0].ne[0] == nil || prevNodes[0].ne[0].v != target {
		return false
	}
	// 从下往上删除
	for i := 0; i < this.lev && prevNodes[i].ne[i] != nil && prevNodes[i].ne[i] == prevNodes[0].ne[0]; i++ {
		prevNodes[i].ne[i] = prevNodes[i].ne[i].ne[i]
	}
	return true
}

// 找到大于等于指定数的前一个节点
func (this *SkipList) find(target int) []*Node {
	prev := make([]*Node, this.lev)
	p := this.head
	for i := this.lev - 1; i >= 0; i-- {
		for p.ne[i] != nil && p.ne[i].v < target {
			p = p.ne[i]
		}
		prev[i] = p
	}
	return prev
}
