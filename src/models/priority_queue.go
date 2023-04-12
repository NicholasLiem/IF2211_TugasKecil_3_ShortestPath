package models

type pqNode[T any] struct {
	value T
	next  *pqNode[T]
}

type number interface {
	int | int8 | int32 | int64 | uint | uint8 | uint32 | uint64 | float32 | float64
}

type PriorityQueue[T any, U number] struct {
	head         *pqNode[T]
	priorityFunc func(T) U
	size         int
}

func NewPriorityQueue[T any, U number](byPriority func(T) U) PriorityQueue[T, U] {
	return PriorityQueue[T, U]{
		head:         nil,
		priorityFunc: byPriority,
		size:         0,
	}
}

func (pq *PriorityQueue[T, U]) Size() int {
	return pq.size
}

func (pq *PriorityQueue[T, U]) IsEmpty() bool {
	return pq.size == 0
}

func (pq *PriorityQueue[T, U]) Enqueue(value T) {
	if pq.IsEmpty() {
		pq.head = &pqNode[T]{value: value, next: nil}
	} else {
		var pred *pqNode[T] = nil
		p := pq.head
		for p != nil {
			if pq.priorityFunc(value) >= pq.priorityFunc(p.value) {
				break
			} else {
				pred = p
				p = p.next
			}
		}
		if pred == nil {
			pq.head = &pqNode[T]{value: value, next: pq.head}
		} else {
			pred.next = &pqNode[T]{value: value, next: p}
		}
	}
	pq.size++
}

func (pq *PriorityQueue[T, U]) Dequeue() T {
	if pq.IsEmpty() {
		panic("Priority queue is empty") // again, this is actually dangerous
	}
	val := pq.head.value
	pq.head = pq.head.next // no need to free memory, thanks to garbage collector
	pq.size--
	return val
}

func (pq *PriorityQueue[T, U]) GetItems() []T {
	items := make([]T, pq.size)
	p := pq.head
	for i := range items {
		items[i] = p.value
		p = p.next
	}
	return items
}
