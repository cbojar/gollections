package collections

import "testing"

func TestImmutableStack(t *testing.T) {
	testStack(t, ImmutableStack())
}

func TestMutableStack(t *testing.T) {
	testStack(t, MutableStack())
}

func testStack(t *testing.T, stack Stack) {
	stack = stack.Push(1)
	if stack.Size() != 1 {
		t.Fatalf("expected stack size 1 but was %d", stack.Size())
	}

	stack = stack.Push(2)
	if stack.Size() != 2 {
		t.Fatalf("expected stack size 2 but was %d", stack.Size())
	}

	value, err := stack.Top()
	if err != nil {
		t.Fatalf("unexpected error getting top value: %v", err)
	}

	if value != 2 {
		t.Fatalf("expected top value 2 but was %d", value)
	}

	stack, err = stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error popping stack: %v", err)
	}

	if stack.Size() != 1 {
		t.Fatalf("expected size to be 1 but was %d", stack.Size())
	}

	value, err = stack.Top()
	if err != nil {
		t.Fatalf("unexpected error getting top value: %v", err)
	}

	if value != 1 {
		t.Fatalf("expected top value 1 but was %d", value)
	}

	stack, err = stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error popping stack: %v", err)
	}

	if stack.Size() != 0 {
		t.Fatalf("expected size to be 0 but was %d", stack.Size())
	}

	_, err = stack.Top()
	if err == nil {
		t.Fatalf("expected error but none returned")
	}

	stack, err = stack.Pop()
	if err == nil {
		t.Fatalf("expected error but none returned")
	}
}

func TestImmutableQueue(t *testing.T) {
	testQueue(t, ImmutableQueue())
}

func TestMutableQueue(t *testing.T) {
	testQueue(t, MutableQueue())
}

func testQueue(t *testing.T, queue Queue) {
	queue = queue.Enqueue(1)
	if queue.Size() != 1 {
		t.Fatalf("expected queue size 1 but was %d", queue.Size())
	}

	queue = queue.Enqueue(2)
	if queue.Size() != 2 {
		t.Fatalf("expected queue size 2 but was %d", queue.Size())
	}

	value, err := queue.Front()
	if err != nil {
		t.Fatalf("unexpected error getting front value: %v", err)
	}

	if value != 1 {
		t.Fatalf("expected front value 1 but was %d", value)
	}

	queue, err = queue.Dequeue()
	if err != nil {
		t.Fatalf("unexpected error dequeueing: %v", err)
	}

	if queue.Size() != 1 {
		t.Fatalf("expected size 1 but was %d", queue.Size())
	}

	value, err = queue.Front()
	if err != nil {
		t.Fatalf("unexpected error getting front value: %v", err)
	}

	if value != 2 {
		t.Fatalf("expected front value 2 but was %d", value)
	}

	queue, err = queue.Dequeue()
	if err != nil {
		t.Fatalf("unexpected error popping stack: %v", err)
	}

	if queue.Size() != 0 {
		t.Fatalf("expected size 0 but was %d", queue.Size())
	}

	_, err = queue.Front()
	if err == nil {
		t.Fatalf("expected error but none returned")
	}

	queue, err = queue.Dequeue()
	if err == nil {
		t.Fatalf("expected error but none returned")
	}
}
