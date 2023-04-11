package models

type AdjacencyMatrix struct {
	Matrix       [][]float64
	ColumnLabels []string
	Latitudes    []float64
	Longitudes   []float64
	NodesCount   int
}

func (am AdjacencyMatrix) GetNodesCount() int {
	return am.NodesCount
}

func NewAdjacencyMatrix(nodeCount int, columnLabels []string, latitudes, longitudes []float64) AdjacencyMatrix {
	matrix := make([][]float64, nodeCount)
	for i := range matrix {
		matrix[i] = make([]float64, nodeCount)
	}

	return AdjacencyMatrix{
		Matrix:       matrix,
		ColumnLabels: columnLabels,
		NodesCount:   nodeCount,
		Latitudes:    latitudes,
		Longitudes:   longitudes,
	}
}
