package test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/NicholasLiem/IF2211_TugasKecil_3_RoutePlanning/models"
	"github.com/NicholasLiem/IF2211_TugasKecil_3_RoutePlanning/utils"
)

func TestAStartSearch(t *testing.T) {
	dir, err := filepath.Abs("./")
	if err != nil {
		t.Errorf(err.Error())
	}
	adjMat, err := utils.AdjacencyMatrixFromFile(filepath.Join(dir, "tc2"))
	if err != nil {
		t.Errorf(err.Error())
	}
	g := *models.NewGraphFromAdjacencyMatrix(*adjMat)
	trace, cost := models.AStarSearch(g, 0, 4)
	traceName := make([]string, len(trace))
	for i, nodeIndex := range trace {
		traceName[i] = g.Nodes[nodeIndex].Name
	}
	fmt.Println("Trace: ", strings.Join(traceName, " -> "))
	fmt.Println("Cost: ", cost)

	expectedTrace := []string{"Bandung", "Purwakarta", "Jakarta"}
	if !reflect.DeepEqual(traceName, expectedTrace) {
		t.Errorf("Expected path %v, but got %v", expectedTrace, traceName)
	}
}
