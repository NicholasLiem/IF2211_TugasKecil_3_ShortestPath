package models

type ucsnode struct {
	nodeIndex int
	g         float64
	trace     []int
}

func UniformCostSearch(graph Graph, src, dest int) ([]int, float64) {
	pq := NewPriorityQueue(func(value ucsnode) float64 {
		return -value.g
	})

	pq.Enqueue(ucsnode{nodeIndex: src, g: 0, trace: []int{src}})

	visited := make(map[int]bool)

	for !pq.IsEmpty() {
		curr := pq.Dequeue()
		if visited[curr.nodeIndex] {
			continue
		}
		visited[curr.nodeIndex] = true

		if curr.nodeIndex == dest {
			return curr.trace, curr.g
		}

		for neighbour, distance := range graph.Edges[curr.nodeIndex] {
			if visited[neighbour] {
				continue
			}
			next := ucsnode{
				nodeIndex: neighbour,
				g:         curr.g + distance,
				trace:     append(append([]int{}, curr.trace...), neighbour),
			}
			pq.Enqueue(next)
		}
	}
	return []int{}, 0
}
