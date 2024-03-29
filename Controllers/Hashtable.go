package Controllers

import "fmt"

const ArraySize = 7

type Hashtable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func Init() *Hashtable {
	hashTable := &Hashtable{}
	for i := range hashTable.array {
		hashTable.array[i] = &bucket{}
	}

	return hashTable
}

func (h *Hashtable) Insert(key string) {
	index := hash(key)
	h.array[index].Insert(key)
}

func (h *Hashtable) Search(key string) bool {
	index := hash(key)
	return h.array[index].Search(key)
}

func (h *Hashtable) Delete(key string) {
	index := hash(key)
	h.array[index].Delete(key)
}

func (b *bucket) Insert(key string) {
	if !b.Search(key) {
		node := &bucketNode{key: key}
		node.next = b.head
		b.head = node
	} else {
		fmt.Println("already exists")
	}
}

func (b *bucket) Delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == key {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

func (b *bucket) Search(key string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func hash(key string) int {
	sum := 0
	for _, val := range key {
		sum += int(val)
	}

	return sum % ArraySize
}

func HashMain() {
	testHash := Init()
	list := []string{
		"hello", "world", "goodbye", "cruel", "world", "yash"}

	for _, val := range list {
		testHash.Insert(val)
	}

	fmt.Println(testHash.Search("world"))
	fmt.Println(testHash.Search("yas"))
	testHash.Delete("world")
	fmt.Println(testHash.Search("world"))
}
