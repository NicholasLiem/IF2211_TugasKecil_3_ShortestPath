package model

type Node struct {
	Name string
}

type Graph struct {
	Nodes map[int]*Node
	Edges map[int]map[int]int64
}

func (g *Graph) AddNode(index int, name string) {
	//	TODO
}

func (g *Graph) AddEdge(fromIndex, toIndex int, weight int64) {
	//	TODO
}

func NewGraphFromAdjacencyMatrix(am AdjacencyMatrix) *Graph {
	g := &Graph{}
	return g
}
