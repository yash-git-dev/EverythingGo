package Controllers

import "fmt"

type NodeT struct {
	children [26]*NodeT
	isEnd    bool
}

type Trie struct {
	root *NodeT
}

func InitTrie() *Trie {
	result := &Trie{root: &NodeT{}}
	return result
}

func TrieMain() {
	testTrie := InitTrie()
	fmt.Println(testTrie.root)
	testTrie.Insert("hello")
	fmt.Println(testTrie.root.children)
	fmt.Println(testTrie.Search("hello"))
}

func (t *Trie) Insert(w string) {
	currentNode := t.root
	for _, val := range w {
		charIndex := val - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &NodeT{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(w string) bool {
	currentNode := t.root
	for _, val := range w {
		charIndex := val - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}

	return currentNode.isEnd
}
