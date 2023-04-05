package model

type adjacencyMatrix struct {
	matrix       [][]int64
	columnLabels []string
	nodesCount   int
}

func (am adjacencyMatrix) getNodesCount() int {
	return am.nodesCount
}

func newAdjacencyMatrix(nodeCount int, columnLabels []string) adjacencyMatrix {
	matrix := make([][]int64, nodeCount)
	for i := range matrix {
		matrix[i] = make([]int64, nodeCount)
	}

	return adjacencyMatrix{
		matrix:       matrix,
		columnLabels: columnLabels,
		nodesCount:   nodeCount,
	}
}
