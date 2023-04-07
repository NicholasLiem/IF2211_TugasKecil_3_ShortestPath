package models

type ucsnode struct {
	nodeIndex int
	g         int64
	trace     []int
}

func UniformCostSearch(graph Graph, src, dest int) ([]int, int64) {
	pq := NewPriorityQueue(func(value ucsnode) int64 {
		return -value.g
	})

	pq.Enqueue(ucsnode{nodeIndex: src, g: 0, trace: []int{src}})

	visited := map[int]bool{}

	for !pq.IsEmpty() {
		curr := pq.Dequeue()
		visited[curr.nodeIndex] = true

		if curr.nodeIndex == dest {
			return curr.trace, int64(curr.g)
		}

		for neighbour, distance := range graph.Edges[curr.nodeIndex] {
			if !visited[neighbour] {
				x := ucsnode{}
				x.nodeIndex = neighbour
				x.g = curr.g + distance
				x.trace = append(curr.trace, neighbour)
				pq.Enqueue(x)
			}
		}
	}
	return []int{}, 0
}
