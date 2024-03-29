package Controllers

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if k > n.Key {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if k < n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}

	if key > n.Key {
		return n.Right.Search(key)
	} else if key < n.Key {
		return n.Left.Search(key)
	}

	return true
}

func CreateTree() {
	tree := Node{Key: 100}
	tree.Insert(50)
	tree.Insert(20)
	tree.Insert(200)
	tree.Insert(120)
	fmt.Println(tree)

	fmt.Println(tree.Search(121))
}
