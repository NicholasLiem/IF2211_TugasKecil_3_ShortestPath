package models

const idxUndef int = -1

type Queue[T interface{}] struct {
	items    []T
	capacity int
	head     int
	tail     int
}

func NewQueue[T interface{}]() Queue[T] {
	return Queue[T]{
		items:    make([]T, 8),
		capacity: 8,
		head:     idxUndef,
		tail:     idxUndef,
	}
}

func NewQueueWithCapacity[T interface{}](capacity int) Queue[T] {
	return Queue[T]{
		items:    make([]T, capacity),
		capacity: capacity,
		head:     idxUndef,
		tail:     idxUndef,
	}
}

func (q Queue[T]) isFull() bool {
	return (q.tail+1)%q.capacity == q.head
}

func (q Queue[T]) IsEmpty() bool {
	return q.head == idxUndef && q.tail == idxUndef
}

func (q Queue[T]) Length() int {
	if q.IsEmpty() {
		return 0
	} else if q.head <= q.tail {
		return q.tail - q.head + 1
	} else {
		return q.capacity - q.head + q.tail + 1
	}
}

func (q *Queue[T]) Enqueue(val T) {
	if q.isFull() {
		q.capacity *= 2
		q.items = append(q.items, make([]T, q.capacity)...)
	}
	if q.IsEmpty() {
		q.head = 0
		q.tail = 0
		q.items[q.head] = val
	} else {
		q.tail = (q.tail + 1) % q.capacity
		q.items[q.tail] = val
	}
}

func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		panic("Queueu is empty") // this is actually a dangerous thing
	}
	val := q.items[q.head]
	if q.tail == q.head {
		q.head = idxUndef
		q.tail = idxUndef
	} else {
		q.head = (q.head + 1) % q.capacity
	}
	return val
}

func (q Queue[T]) GetItems() []T {
	if q.IsEmpty() {
		return []T{}
	}
	arr := make([]T, q.Length())
	for i := range arr {
		arr[i] = q.items[(q.head+i)%q.capacity]
	}
	return arr
}
