package test

import (
	"main/models"
	"reflect"
	"testing"
)

func TestNewGraphFromAdjacencyMatrix(t *testing.T) {
	columnLabels := []string{"A", "B", "C", "D"}
	nodeCount := len(columnLabels)
	adjacencyMatrix := models.NewAdjacencyMatrix(nodeCount, columnLabels)

	adjacencyMatrix.Matrix[0][0] = 0
	adjacencyMatrix.Matrix[0][1] = 1
	adjacencyMatrix.Matrix[0][2] = 2
	adjacencyMatrix.Matrix[0][3] = 0

	adjacencyMatrix.Matrix[1][0] = 1
	adjacencyMatrix.Matrix[1][1] = 0
	adjacencyMatrix.Matrix[1][2] = 3
	adjacencyMatrix.Matrix[1][3] = 0

	adjacencyMatrix.Matrix[2][0] = 2
	adjacencyMatrix.Matrix[2][1] = 3
	adjacencyMatrix.Matrix[2][2] = 0
	adjacencyMatrix.Matrix[2][3] = 4

	adjacencyMatrix.Matrix[3][0] = 0
	adjacencyMatrix.Matrix[3][1] = 0
	adjacencyMatrix.Matrix[3][2] = 4
	adjacencyMatrix.Matrix[3][3] = 0

	graph := models.NewGraphFromAdjacencyMatrix(adjacencyMatrix)

	expectedNodeNames := map[int]string{0: "A", 1: "B", 2: "C", 3: "D"}
	for i, node := range graph.Nodes {
		if node.Name != expectedNodeNames[i] {
			t.Errorf("Expected node name %s but got %s", expectedNodeNames[i], node.Name)
		}
	}

	expectedEdges := map[int]map[int]int64{
		0: {1: 1, 2: 2},
		1: {0: 1, 2: 3},
		2: {0: 2, 1: 3, 3: 4},
		3: {2: 4},
	}

	if !reflect.DeepEqual(graph.Edges, expectedEdges) {
		t.Errorf("Expected edges %v but got %v", expectedEdges, graph.Edges)
	}
}
