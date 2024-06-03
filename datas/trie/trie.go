package trie

// 字典树
type Trie struct {
	root *node
}

type node struct {
	exist bool
	ne    map[byte]*node
}

func NewTrie() Trie {
	return Trie{
		root: &node{
			exist: false,
			ne:    make(map[byte]*node, 26),
		},
	}
}

func (tr *Trie) Insert(word string) {
	pc := tr.root
	for i := range word {
		if pc.ne[word[i]] == nil {
			pc.ne[word[i]] = new(node)
			pc.ne[word[i]].ne = make(map[byte]*node)
		}
		pc = pc.ne[word[i]]
	}
	pc.exist = true
}

func (tr *Trie) Search(word string) bool {
	pc := tr.root
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

func (tr *Trie) StartsWith(prefix string) bool {
	pc := tr.root
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
