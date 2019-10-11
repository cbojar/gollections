package collections

import (
	"fmt"
	"strings"
)

// MutableStack creates an empty instance of a stack implemented mutably
func MutableStack() Stack {
	values := make([]interface{}, 1)
	return &mutableStack{values: values, head: 0}
}

type mutableStack struct {
	values []interface{}
	head   int
}

func (s *mutableStack) Push(i interface{}) Stack {
	if s.head >= len(s.values) {
		s.grow()
	}

	s.values[s.head] = i
	s.head++

	return s
}

func (s *mutableStack) Pop() (Stack, error) {
	if s.head == 0 {
		return nil, fmt.Errorf("cannot pop empty stack")
	}

	s.head--
	s.values[s.head] = nil

	if s.head < len(s.values)/2 {
		s.shrink()
	}

	return s, nil
}

func (s *mutableStack) Top() (interface{}, error) {
	if s.head == 0 {
		return nil, fmt.Errorf("cannot get top value of empty stack")
	}

	return s.values[s.head-1], nil
}

func (s *mutableStack) Size() int {
	return s.head
}

func (s *mutableStack) Capacity() int {
	return len(s.values)
}

func (s *mutableStack) String() string {
	var output strings.Builder

	fmt.Fprint(&output, "<->[")
	for i := s.head - 1; i >= 0; i-- {
		separator := ", "
		if i == 0 {
			separator = ""
		}
		fmt.Fprintf(&output, "%v%s", s.values[i], separator)
	}
	fmt.Fprint(&output, "[")

	return output.String()
}

func (s *mutableStack) grow() {
	oldValues := s.values
	s.values = make([]interface{}, len(oldValues)*2)

	copy(s.values, oldValues)
}

func (s *mutableStack) shrink() {
	oldValues := s.values
	s.values = make([]interface{}, len(oldValues)/2)

	copy(s.values, oldValues)
}
