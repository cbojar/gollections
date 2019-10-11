package collections

import (
	"fmt"
	"strings"
)

var emptyStackInstance = emptyStack{}

// ImmutableStack creates an empty instance of a stack implemented immutably
func ImmutableStack() Stack {
	return &emptyStackInstance
}

type valueStack struct {
	previous Stack
	value    interface{}
	depth    int
}

func (s *valueStack) Push(i interface{}) Stack {
	return &valueStack{previous: s, value: i, depth: s.Size() + 1}
}

func (s *valueStack) Pop() (Stack, error) {
	return s.previous, nil
}

func (s *valueStack) Top() (interface{}, error) {
	return s.value, nil
}

func (s *valueStack) Size() int {
	return s.depth
}

func (s *valueStack) Capacity() int {
	return s.Size()
}

func (s *valueStack) String() string {
	var output strings.Builder

	fmt.Fprint(&output, "<->[")
	var current Stack = s
	for ; current.Size() > 0; current, _ = current.Pop() {
		separator := ", "
		if current.Size() == 1 {
			separator = ""
		}
		top, _ := current.Top()
		fmt.Fprintf(&output, "%v%s", top, separator)
	}
	fmt.Fprint(&output, "[")

	return output.String()
}

type emptyStack struct{}

func (s *emptyStack) Push(i interface{}) Stack {
	return &valueStack{previous: s, value: i, depth: 1}
}

func (s *emptyStack) Pop() (Stack, error) {
	return nil, fmt.Errorf("cannot pop empty stack")
}

func (s *emptyStack) Top() (interface{}, error) {
	return nil, fmt.Errorf("cannot get top value of empty stack")
}

func (s *emptyStack) Size() int {
	return 0
}

func (s *emptyStack) Capacity() int {
	return s.Size()
}

func (s *emptyStack) String() string {
	return "<->[["
}
