package test

import (
	"github.com/NicholasLiem/IF2211_TugasKecil_3_RoutePlanning/models"
	"github.com/NicholasLiem/IF2211_TugasKecil_3_RoutePlanning/utils"
	"path/filepath"
	"reflect"
	"testing"
)

func TestNewGraphFromAdjacencyMatrix(t *testing.T) {
	columnLabels := []string{"A", "B", "C", "D"}
	latitudes := []float64{13.115, 2.8, 5, 9.25}
	longitudes := []float64{3.5, 7.45, 12.35, 15.8}
	nodeCount := len(columnLabels)
	adjacencyMatrix := models.NewAdjacencyMatrix(nodeCount, columnLabels, latitudes, longitudes)

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
	expectedNodeLatitudes := map[int]float64{0: 13.115, 1: 2.8, 2: 5, 3: 9.25}
	expectedNodeLongitudes := map[int]float64{0: 3.5, 1: 7.45, 2: 12.35, 3: 15.8}
	for i, node := range graph.Nodes {
		if node.Name != expectedNodeNames[i] {
			t.Errorf("Expected node name %s but got %s", expectedNodeNames[i], node.Name)
		}
		if node.Latitude != expectedNodeLatitudes[i] {
			t.Errorf("Expected node latitude %f but got %f", expectedNodeLatitudes[i], node.Latitude)
		}
		if node.Longitude != expectedNodeLongitudes[i] {
			t.Errorf("Expected node longitude %f but got %f", expectedNodeLongitudes[i], node.Longitude)
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

func TestAdjacencyMatrixFromFile(t *testing.T) {
	dir, err := filepath.Abs("./")
	if err != nil {
		t.Errorf("[ERROR] cannot get source file path")
	}
	adjMat, err := utils.AdjacencyMatrixFromFile(dir + "/tc1.txt")
	if err != nil {
		t.Errorf(err.Error())
	}
	expectedNodesCount := 3
	expectedColumLabels := []string{"A", "B", "C"}
	expectedLatitudes := []float64{1, 2, 3}
	expectedLongitudes := []float64{3, 2, 1}
	if adjMat.NodesCount != expectedNodesCount {
		t.Errorf("Expected nodes count %d but got %d", expectedNodesCount, adjMat.NodesCount)
	}
	if !reflect.DeepEqual(adjMat.ColumnLabels, expectedColumLabels) {
		t.Errorf("Expected column labels %v but got %v", expectedColumLabels, adjMat.ColumnLabels)
	}
	if !reflect.DeepEqual(adjMat.Latitudes, expectedLatitudes) {
		t.Errorf("Expected latitudes %v but got %v", expectedLatitudes, adjMat.Latitudes)
	}
	if !reflect.DeepEqual(adjMat.Longitudes, expectedLongitudes) {
		t.Errorf("Expected longitudes %v but got %v", expectedLongitudes, adjMat.Longitudes)
	}
}
