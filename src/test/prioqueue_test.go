package test

import (
	"github.com/NicholasLiem/IF2211_TugasKecil_3_RoutePlanning/models"
	"testing"
)

func byPriority(node models.Node) int {
	return len(node.Name)
}

func TestPriorityQueue(t *testing.T) {
	pq := models.NewPriorityQueue(byPriority)

	node1 := models.Node{Name: "Arad"}
	node2 := models.Node{Name: "Bucharest"}
	node3 := models.Node{Name: "Sibiu"}
	node4 := models.Node{Name: "Pitesti"}

	pq.Enqueue(node1)
	pq.Enqueue(node2)
	pq.Enqueue(node3)
	pq.Enqueue(node4)

	expectedOrder := []models.Node{node2, node4, node3, node1}

	for _, expectedItem := range expectedOrder {
		item := pq.Dequeue()

		if item.Name != expectedItem.Name {
			t.Errorf("Expected node name %s, but got %s", expectedItem.Name, item.Name)
		}
	}
}
