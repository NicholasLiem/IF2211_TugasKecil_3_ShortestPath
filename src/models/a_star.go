package models

import (
	"math"
)

type astarnode struct {
	nodeIndex int
	f         float64
	g         float64
	trace     []int
}

func distance(lat1, lon1, lat2, lon2 float64) float64 {
	p := 0.017453292519943295 // Math.PI / 180
	c := math.Cos
	a := 0.5 - c((lat2-lat1)*p)/2 +
		c(lat1*p)*c(lat2*p)*
			(1-c((lon2-lon1)*p))/2

	return 12742 * math.Asin(math.Sqrt(a)) // 2 * R; R = 6371 km
}

func calculateH(graph Graph, current, destination int) float64 {
	node1 := graph.Nodes[current]
	node2 := graph.Nodes[destination]

	lat1 := node1.Latitude
	lat2 := node2.Latitude
	lon1 := node1.Longitude
	lon2 := node2.Longitude

	return distance(lat1, lon1, lat2, lon2)
}

func AStarSearch(graph Graph, src, dest int) ([]int, int64) {
	pq := NewPriorityQueue(func(value astarnode) float64 {
		return -value.f
	})

	pq.Enqueue(astarnode{nodeIndex: src, f: 0, g: 0, trace: []int{src}})

	visited := map[int]bool{}

	for !pq.IsEmpty() {
		curr := pq.Dequeue()
		visited[curr.nodeIndex] = true

		if curr.nodeIndex == dest {
			return curr.trace, int64(curr.g)
		}

		for neighbour, distance := range graph.Edges[curr.nodeIndex] {
			if !visited[neighbour] {
				x := astarnode{}
				x.nodeIndex = neighbour
				x.g = curr.g + float64(distance)
				x.f = x.g + calculateH(graph, curr.nodeIndex, dest)
				x.trace = append(curr.trace, neighbour)
				pq.Enqueue(x)
			}
		}
	}
	return []int{}, 0
}
