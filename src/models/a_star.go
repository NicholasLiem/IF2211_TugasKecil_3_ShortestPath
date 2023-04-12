package models

import (
	"github.com/NicholasLiem/IF2211_TugasKecil_3_RoutePlanning/utils"
)

type astarnode struct {
	nodeIndex int
	f         float64
	g         float64
	trace     []int
}

func calculateH(graph Graph, current, destination int) float64 {
	node1 := graph.Nodes[current]
	node2 := graph.Nodes[destination]

	lat1 := node1.Latitude
	lat2 := node2.Latitude
	lon1 := node1.Longitude
	lon2 := node2.Longitude

	return utils.Distance(lat1, lon1, lat2, lon2)
}

func AStarSearch(graph Graph, src, dest int) ([]int, float64) {
	pq := NewPriorityQueue(func(value astarnode) float64 {
		return -value.f
	})

	pq.Enqueue(astarnode{nodeIndex: src, f: 0, g: 0, trace: []int{src}})

	visited := map[int]bool{}

	for !pq.IsEmpty() {
		curr := pq.Dequeue()
		visited[curr.nodeIndex] = true

		if curr.nodeIndex == dest {
			return curr.trace, float64(curr.g)
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
