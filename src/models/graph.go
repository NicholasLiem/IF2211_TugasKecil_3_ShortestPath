package model

type node struct {
	name string
}

type graph struct {
	nodes map[int]*node
	edges map[int]map[int]int64
}

func (g *graph) addNode(index int, name string) {
	//	TODO
}

func (g *graph) addEdge(fromIndex, toIndex int, weight int64) {
	//	TODO
}

func newGraphFromAdjacencyMatrix(am adjacencyMatrix) *graph {
	g := &graph{}
	return g
}
