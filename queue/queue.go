package queue

import (
	"github.com/iAziz786/algosandds/linkedlist2"
)

type Q struct {
	items *linkedlist2.LinkedList
	head  *int
	tail  *int
}

func New() *Q {
	return &Q{
		items: linkedlist2.New(),
	}
}

func (q *Q) Enqueue(val int) {
	// TODO(aziz): Use tail pointer for O(1) insertion
	q.items.PushBack(val)
}

func (q *Q) Dequeue() *int {
	res, err := q.items.ValueAt(0)
	if err != nil {
		return nil
	}
	q.items.PopFront()
	return &res
}

func (q Q) Empty() bool {
	return false
}
