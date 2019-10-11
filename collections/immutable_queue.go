package collections

import (
	"fmt"
	"strings"
)

// ImmutableQueue creates an empty instance of a queue implemented immutably
func ImmutableQueue() Queue {
	return &immutableQueue{stack: ImmutableStack()}
}

type immutableQueue struct {
	stack Stack
}

func (q *immutableQueue) Enqueue(i interface{}) Queue {
	return &immutableQueue{stack: q.stack.Push(i)}
}

func (q *immutableQueue) Dequeue() (Queue, error) {
	if q.Size() == 0 {
		return nil, fmt.Errorf("Cannot dequeue an empty queue")
	}

	reverse := reverseStack(q.stack, ImmutableStack())
	reverse, _ = reverse.Pop()
	forward := reverseStack(reverse, ImmutableStack())

	return &immutableQueue{stack: forward}, nil
}

func (q *immutableQueue) Front() (interface{}, error) {
	if q.Size() == 0 {
		return nil, fmt.Errorf("cannot get front value of empty queue")
	}

	current := q.stack
	for current.Size() > 1 {
		current, _ = current.Pop()
	}

	top, _ := current.Top()
	return top, nil
}

func (q *immutableQueue) Size() int {
	return q.stack.Size()
}

func (q *immutableQueue) Capacity() int {
	return q.stack.Capacity()
}

func (q *immutableQueue) String() string {
	var output strings.Builder

	fmt.Fprint(&output, "<-[")
	for current := q.stack; current.Size() > 0; current, _ = current.Pop() {
		separator := ", "
		if current.Size() == 1 {
			separator = ""
		}
		top, _ := current.Top()
		fmt.Fprintf(&output, "%v%s", top, separator)
	}
	fmt.Fprint(&output, "[<-")

	return output.String()
}
