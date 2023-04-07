package models

type pqNode[T interface{}] struct {
	value T
	next  *pqNode[T]
}

type PriorityQueue[T interface{}] struct {
	head         *pqNode[T]
	priorityFunc func(T) int
	size         int
}

func NewPriorityQueue[T interface{}](byPriority func(T) int) PriorityQueue[T] {
	return PriorityQueue[T]{
		head:         nil,
		priorityFunc: byPriority,
		size:         0,
	}
}

func (pq PriorityQueue[T]) Size() int {
	return pq.size
}

func (pq PriorityQueue[T]) IsEmpty() bool {
	return pq.size == 0
}

func (pq *PriorityQueue[T]) Enqueue(value T) {
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

func (pq *PriorityQueue[T]) Dequeue() T {
	if pq.IsEmpty() {
		panic("Priority queue is empty") // again, this is actually dangerous
	}
	val := pq.head.value
	pq.head = pq.head.next // no need to free memory, thanks to garbage collector
	pq.size--
	return val
}

func (pq PriorityQueue[T]) GetItems() []T {
	items := make([]T, pq.size)
	p := pq.head
	for i := range items {
		items[i] = p.value
		p = p.next
	}
	return items
}
