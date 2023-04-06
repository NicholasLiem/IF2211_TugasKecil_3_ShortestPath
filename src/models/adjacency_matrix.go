package models

type AdjacencyMatrix struct {
	Matrix       [][]int64
	ColumnLabels []string
	Latitudes    []float64
	Longitudes   []float64
	NodesCount   int
}

func (am AdjacencyMatrix) GetNodesCount() int {
	return am.NodesCount
}

func NewAdjacencyMatrix(nodeCount int, columnLabels []string, latitudes, longitudes []float64) AdjacencyMatrix {
	matrix := make([][]int64, nodeCount)
	for i := range matrix {
		matrix[i] = make([]int64, nodeCount)
	}

	return AdjacencyMatrix{
		Matrix:       matrix,
		ColumnLabels: columnLabels,
		NodesCount:   nodeCount,
		Latitudes:    latitudes[:len(columnLabels)],
		Longitudes:   longitudes,
	}
}
