package model

type AdjacencyMatrix struct {
	Matrix       [][]int64
	ColumnLabels []string
	NodesCount   int
}

func (am AdjacencyMatrix) GetNodesCount() int {
	return am.NodesCount
}

func NewAdjacencyMatrix(nodeCount int, columnLabels []string) AdjacencyMatrix {
	matrix := make([][]int64, nodeCount)
	for i := range matrix {
		matrix[i] = make([]int64, nodeCount)
	}

	return AdjacencyMatrix{
		Matrix:       matrix,
		ColumnLabels: columnLabels,
		NodesCount:   nodeCount,
	}
}
