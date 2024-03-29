package Controllers

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head   *node
	length int
}

func (l *LinkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func LinkMain() {
	mylist := LinkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 15}
	node3 := &node{data: 3}

	mylist.prepend(node3)
	mylist.prepend(node2)
	mylist.prepend(node1)

	mylist.printData()
	mylist.remove(48)
	mylist.printData()

}

func (l *LinkedList) remove(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	toDelete := l.head
	for toDelete.next.data != value {
		if toDelete.next.next == nil {
			return
		}

		toDelete = toDelete.next
	}
	toDelete.next = toDelete.next.next
	l.length--
}

func (l LinkedList) printData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println("")
}
