package Controllers

import "fmt"

type MaxHeap struct {
	data []int
}

func (h *MaxHeap) Insert(key int) {
	h.data = append(h.data, key)
	h.MaxHeapifyUp(len(h.data) - 1)
}

func (h *MaxHeap) MaxHeapifyUp(n int) {
	for h.data[parent(n)] < h.data[n] {
		h.Swap(parent(n), n)
		n = parent(n)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) Swap(i1, i2 int) {
	h.data[i1], h.data[i2] = h.data[i2], h.data[i1]
}

func (h *MaxHeap) Extract() int {
	extracted := h.data[0]
	l := len(h.data) - 1

	if len(h.data) == 0 {
		fmt.Println("empty heap")
		return -1
	}

	h.data[0] = h.data[l]
	h.data = h.data[:l]

	h.MaxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) MaxHeapifyDown(n int) {
	lastIndex := len(h.data) - 1

	l, r := left(n), right(n)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.data[l] > h.data[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.data[n] > h.data[childToCompare] {
			h.Swap(n, childToCompare)
			n = childToCompare
			l, r = left(n), right(n)
		} else {
			return
		}
	}
}

func HeapMain() {
	buildHeap := []int{10, 20, 5, 33, 99, 48}

	maxHeap := &MaxHeap{}

	for _, val := range buildHeap {
		maxHeap.Insert(val)
	}

	fmt.Println(maxHeap)

	for i := 0; i < 5; i++ {
		maxHeap.Extract()
		fmt.Println(maxHeap)
	}
}
