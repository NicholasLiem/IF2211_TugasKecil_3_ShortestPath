package test

import (
	"reflect"
	"testing"

	"github.com/NicholasLiem/IF2211_TugasKecil_3_RoutePlanning/models"
)

func TestQueue(t *testing.T) {
	q := models.NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	q.Enqueue(7)
	q.Enqueue(8)
	q.Enqueue(9)

	expectedValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(expectedValues, q.GetItems()) {
		t.Errorf("Expected values %v, but got %v", expectedValues, q.GetItems())
	}

	actualValues := []int{q.Dequeue(), q.Dequeue(), q.Dequeue()}
	if q.IsEmpty() {
		t.Errorf("Expected queue to be not empty")
	}

	expectedValues = []int{1, 2, 3}
	if !reflect.DeepEqual(actualValues, expectedValues) {
		t.Errorf("Expected values %v, but got %v", expectedValues, actualValues)
	}

	expectedValues = []int{4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(expectedValues, q.GetItems()) {
		t.Errorf("Expected values %v, but got %v", expectedValues, q.GetItems())
	}
}
