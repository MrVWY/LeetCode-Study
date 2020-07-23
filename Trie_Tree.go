package main

import "fmt"

//字典树
type Trie struct {
	Root *TrieNode
}

type TrieNode struct {
	Children map[string]*TrieNode
	isEnd bool
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		Children: make(map[string]*TrieNode),
		isEnd:    false,
	}
}

func newTrieRoot() *Trie {
	return &Trie{Root: newTrieNode()}
}

func (t *Trie) Insert(words []string) {
	//注意如果是以数组形式传进来, 记得重置根节点
	for _, word := range words{
		root := t.Root //重置根节点
		for i := 0 ; i < len(word) ; i++ {
			str := string(word[i])
			if _, ok := root.Children[str] ; !ok {
				root.Children[str] = newTrieNode()
			}
			root = root.Children[str]
			root.isEnd = true
		}
	}
}

//Search word or prefix
func (t *Trie) Search(word string) bool {
	root := t.Root
	for i := 0 ; i < len(word) ; i++ {
		str := string(word[i])
		if _, ok := root.Children[str] ; !ok {
			return false
		}
		root = root.Children[str]
	}
	return true
}

func main() {
	r := newTrieRoot()
	r.Insert([]string{"ab","abc","bc"})

	fmt.Println(r.Search("bc"))
}