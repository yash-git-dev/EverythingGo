package Controllers

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	l := len(q.items)
	if l == 0 {
		return 0
	}
	firstItem := q.items[0]
	q.items = q.items[1:]
	return firstItem
}
