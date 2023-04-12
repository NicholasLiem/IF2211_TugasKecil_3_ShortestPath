package models

import (
	"github.com/NicholasLiem/Tucil3_13521083_13521135/utils"
	"math"
)

type astarnode struct {
	nodeIndex int
	f         float64
	g         float64
	h         float64
	trace     []int
}

func calculateH(graph Graph, current, destination int, weightRange float64) float64 {
	node1 := graph.Nodes[current]
	node2 := graph.Nodes[destination]

	lat1 := node1.Latitude
	lat2 := node2.Latitude
	lon1 := node1.Longitude
	lon2 := node2.Longitude

	return utils.Distance(lat1, lon1, lat2, lon2, weightRange)
}

func getWeightRange(graph Graph) float64 {
	min := math.MaxFloat64
	max := 0.0
	for from := range graph.Edges {
		for _, weight := range graph.Edges[from] {
			if weight < min {
				min = weight
			} else if weight > max {
				max = weight
			}
		}
	}
	return max - min
}

func AStarSearch(graph Graph, src, dest int) ([]int, float64) {
	pq := NewPriorityQueue(func(value astarnode) float64 {
		return -value.f
	})
	weighRange := getWeightRange(graph)

	pq.Enqueue(astarnode{nodeIndex: src, f: 0, g: 0, trace: []int{src}})

	visited := map[int]bool{}

	for !pq.IsEmpty() {
		curr := pq.Dequeue()
		if visited[curr.nodeIndex] {
			continue
		}
		if curr.nodeIndex == dest {
			return curr.trace, curr.g
		}

		for neighbour, distance := range graph.Edges[curr.nodeIndex] {
			if visited[neighbour] {
				continue
			}
			x := astarnode{
				nodeIndex: neighbour,
				g:         curr.g + distance,
				h:         calculateH(graph, curr.nodeIndex, dest, weighRange),
				trace:     append(append([]int{}, curr.trace...), neighbour),
			}
			x.f = x.g + x.h
			pq.Enqueue(x)
		}
	}
	return []int{}, 0
}
