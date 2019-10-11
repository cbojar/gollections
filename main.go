package main

import (
	"fmt"

	"github.com/cbojar/gollections/collections"
)

func main() {
	runStack(collections.ImmutableStack(), "Immutable Stack")
	runStack(collections.MutableStack(), "Mutable Stack")
	runQueue(collections.ImmutableQueue(), "Immutable Queue")
	runQueue(collections.MutableQueue(), "Mutable Queue")
}

func runStack(stack collections.Stack, description string) {
	fmt.Println("Stack type:", description)

	for i := 10; i >= 1; i-- {
		stack = stack.Push(i)
		top, _ := stack.Top()
		fmt.Printf("Value: %2d, Size: %2d/%2d, Contents: %s\n", top, stack.Size(), stack.Capacity(), stack)
	}

	for stack.Size() > 0 {
		top, _ := stack.Top()
		fmt.Printf("Value: %2d, Size: %2d/%2d, Contents: %s\n", top, stack.Size(), stack.Capacity(), stack)
		stack, _ = stack.Pop()
	}

	fmt.Printf("Size: %2d/%2d, Contents: %s\n", stack.Size(), stack.Capacity(), stack)
}

func runQueue(queue collections.Queue, description string) {
	fmt.Println("Queue type:", description)

	for i := 1; i <= 10; i++ {
		queue = queue.Enqueue(i)
		front, _ := queue.Front()
		fmt.Printf("Value: %2d, Size: %2d/%2d, Contents: %s\n", front, queue.Size(), queue.Capacity(), queue)
	}

	for queue.Size() > 0 {
		front, _ := queue.Front()
		fmt.Printf("Value: %2d, Size: %2d/%2d, Contents: %s\n", front, queue.Size(), queue.Capacity(), queue)
		queue, _ = queue.Dequeue()
	}

	fmt.Printf("Size: %2d/%2d, Contents: %s\n", queue.Size(), queue.Capacity(), queue)
}
