package trie

// 字典树
type Trie struct {
	root *Node
}

type Node struct {
	exist bool
	ne    map[byte]*Node
}

func NewTrie() Trie {
	return Trie{
		root: &Node{
			exist: false,
			ne:    make(map[byte]*Node, 26),
		},
	}
}

func (this *Trie) Insert(word string) {
	pc := this.root
	for i := range word {
		if pc.ne[word[i]] == nil {
			pc.ne[word[i]] = new(Node)
			pc.ne[word[i]].ne = make(map[byte]*Node)
		}
		pc = pc.ne[word[i]]
	}
	pc.exist = true
}

func (this *Trie) Search(word string) bool {
	pc := this.root
	for i := range word {
		if pc.ne[word[i]] == nil {
			return false
		}
		if _, ok := pc.ne[word[i]]; !ok {
			return false
		} else {
			pc = pc.ne[word[i]]
		}
	}
	if !pc.exist {
		return false
	}
	return true
}

func (this *Trie) StartsWith(prefix string) bool {
	pc := this.root
	for i := range prefix {
		if pc.ne[prefix[i]] == nil {
			return false
		}
		if _, ok := pc.ne[prefix[i]]; !ok {
			return false
		} else {
			pc = pc.ne[prefix[i]]
		}
	}
	return true
}
