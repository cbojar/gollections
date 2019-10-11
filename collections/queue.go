package collections

// Queue represents a FIFO data structure
type Queue interface {
	Enqueue(i interface{}) Queue
	Dequeue() (Queue, error)
	Front() (interface{}, error)
	Size() int
	Capacity() int
}
