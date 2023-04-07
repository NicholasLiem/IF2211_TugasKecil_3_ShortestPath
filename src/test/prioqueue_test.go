package test

import (
	"container/heap"
	"github.com/NicholasLiem/IF2211_TugasKecil_3_RoutePlanning/models"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := models.PriorityQueue{}

	item1 := &models.Item{Value: &models.Node{Name: "Node 1"}, Priority: 3}
	heap.Push(&pq, item1)

	item2 := &models.Item{Value: &models.Node{Name: "Node 2"}, Priority: 1}
	heap.Push(&pq, item2)

	item3 := &models.Item{Value: &models.Node{Name: "Node 3"}, Priority: 2}
	heap.Push(&pq, item3)

	expectedOrder := []*models.Item{item2, item3, item1}

	for _, expectedItem := range expectedOrder {
		item := heap.Pop(&pq).(*models.Item)

		if item.Priority != expectedItem.Priority {
			t.Errorf("Expected item with priority %d, but got %d", expectedItem.Priority, item.Priority)
		}

		if item.Value.Name != expectedItem.Value.Name {
			t.Errorf("Expected item with value %s, but got %s", expectedItem.Value.Name, item.Value.Name)
		}
	}
}
