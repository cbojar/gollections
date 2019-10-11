package collections

import (
	"fmt"
	"strings"
)

// MutableQueue creates an empty instance of a queue implemented mutably
func MutableQueue() Queue {
	values := make([]interface{}, 1)
	return &mutableQueue{values: values, head: 0, tail: 0}
}

type mutableQueue struct {
	values []interface{}
	head   int
	tail   int
}

func (q *mutableQueue) Enqueue(i interface{}) Queue {
	if q.tail == len(q.values) {
		q.grow()
	}

	q.values[q.tail] = i
	q.tail++

	return q
}

func (q *mutableQueue) Dequeue() (Queue, error) {
	if q.Size() == 0 {
		return nil, fmt.Errorf("cannot dequeue from empty queue")
	}

	q.values[q.head] = nil
	q.head++

	if q.Size() < len(q.values)/2 {
		q.shrink()
	}

	return q, nil
}

func (q *mutableQueue) Front() (interface{}, error) {
	if q.Size() == 0 {
		return nil, fmt.Errorf("cannot get front value of empty queue")
	}

	return q.values[q.head], nil
}

func (q *mutableQueue) Size() int {
	return q.tail - q.head
}

func (q *mutableQueue) Capacity() int {
	return len(q.values)
}

func (q *mutableQueue) String() string {
	var output strings.Builder

	fmt.Fprint(&output, "<-[")
	for i := q.head; i < q.tail; i++ {
		separator := ", "
		if i == q.tail-1 {
			separator = ""
		}
		fmt.Fprintf(&output, "%v%s", q.values[i], separator)
	}
	fmt.Fprint(&output, "[<-")

	return output.String()
}

func (q *mutableQueue) grow() {
	oldValues := q.values
	q.values = make([]interface{}, len(oldValues)*2)

	copy(q.values, oldValues)
}

func (q *mutableQueue) shrink() {
	newCapacity := len(q.values)
	size := q.Size()
	for size < newCapacity/2 {
		newCapacity = newCapacity / 2
	}

	oldValues := q.values
	q.values = make([]interface{}, newCapacity)

	copy(q.values, oldValues[q.head:q.tail])
	q.tail = q.tail - q.head
	q.head = 0
}
