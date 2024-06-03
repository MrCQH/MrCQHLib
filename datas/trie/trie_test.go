package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("abc")
	trie.Insert("acb")
	if !trie.Search("abc") {
		t.Error("abc exist", trie.Search("abc"))
	}
	fmt.Println(trie.StartsWith("db"))
}
