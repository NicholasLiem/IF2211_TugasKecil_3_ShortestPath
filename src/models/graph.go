package models

type Node struct {
	Index     int
	Name      string
	Latitude  float64
	Longitude float64
}

type Graph struct {
	Nodes map[int]*Node
	Edges map[int]map[int]float64
}

func (g *Graph) AddNode(index int, name string, latitude, longitude float64) {
	if g.Nodes == nil {
		g.Nodes = make(map[int]*Node)
	}
	newNode := Node{Index: index, Name: name, Latitude: latitude, Longitude: longitude}
	g.Nodes[index] = &newNode
}

func (g *Graph) AddEdge(fromIndex, toIndex int, weight float64) {
	if g.Edges == nil {
		g.Edges = make(map[int]map[int]float64)
	}
	if _, ok := g.Edges[toIndex]; !ok {
		g.Edges[toIndex] = make(map[int]float64)
	}
	g.Edges[toIndex][fromIndex] = weight
}

func NewGraphFromAdjacencyMatrix(am AdjacencyMatrix) *Graph {
	g := &Graph{
		Nodes: make(map[int]*Node),
		Edges: make(map[int]map[int]float64),
	}
	for i := 0; i < am.NodesCount; i++ {
		g.AddNode(i, am.ColumnLabels[i], am.Latitudes[i], am.Longitudes[i])
	}

	for i := 0; i < am.NodesCount; i++ {
		for j := 0; j < i; j++ {
			if am.Matrix[i][j] == 0 {
				continue
			} else {
				g.AddEdge(i, j, am.Matrix[i][j])
				g.AddEdge(j, i, am.Matrix[i][j])
			}
		}
	}
	return g
}

func (g *Graph) GetEdgeWeight(indexSource, indexDestination int) float64 {
	if _, ok := g.Edges[indexSource]; !ok {
		return -1
	}
	if weight, ok := g.Edges[indexSource][indexDestination]; ok {
		return weight
	}
	return -1
}

func (g *Graph) GetNeighbors(index int) []int {
	if _, ok := g.Nodes[index]; !ok {
		return []int{}
	}
	var neighbors []int
	if edges, ok := g.Edges[index]; ok {
		for nodeIndex := range edges {
			neighbors = append(neighbors, nodeIndex)
		}
	}
	return neighbors
}
