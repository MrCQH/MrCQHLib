package skip_list

import (
	"testing"
)

func TestSkipList(t *testing.T) {
	skiplist := New()
	skiplist.Add(1)
	skiplist.Add(2)
	skiplist.Add(3)
	println(skiplist.Search(0))
	skiplist.Add(4)
	println(skiplist.Search(1))
	println(skiplist.Search(5))
	println(skiplist.Search(3))
	println(skiplist.Search(6))
	skiplist.Erase(1)
	println(skiplist.Search(1))
}
