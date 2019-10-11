package collections

// Stack represents a LIFO data structure
type Stack interface {
	Push(i interface{}) Stack
	Pop() (Stack, error)
	Top() (interface{}, error)
	Size() int
	Capacity() int
}

func reverseStack(from Stack, to Stack) Stack {
	for current := from; current.Size() > 0; current, _ = current.Pop() {
		value, _ := current.Top()
		to = to.Push(value)
	}

	return to
}
