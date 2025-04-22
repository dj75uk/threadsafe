package threadsafe

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	obj := NewQueue[int]()
	if obj == nil {
		t.Errorf("NewQueue[int]() returned nil")
	}
}

func TestQueue_Enqueue(t *testing.T) {
	obj := NewQueue[int]()
	previousLength := obj.Length()
	obj.Enqueue(1)
	currentLength := obj.Length()

	if currentLength != (previousLength + 1) {
		t.Errorf("Enqueue() failed to enqueue element")
	}
}

func TestQueue_TryDequeue_Returns_False_On_Empty_Queue(t *testing.T) {
	obj := NewQueue[int]()

	temp := 0
	actual := obj.TryDequeue(&temp)

	if actual != false {
		t.Errorf("TryDequeue failed to return false")
	}
}

func TestQueue_TryDequeue_Returns_True_On_NonEmpty_Queue(t *testing.T) {
	obj := NewQueue[int]()
	obj.Enqueue(1)

	temp := 0
	actual := obj.TryDequeue(&temp)

	if actual != true {
		t.Errorf("TryDequeue failed to return true")
	}
}

func TestQueue_TryDequeue_Removes_Items(t *testing.T) {
	obj := NewQueue[int]()
	obj.Enqueue(1)

	temp := 0
	_ = obj.TryDequeue(&temp)

	actual := obj.TryDequeue(&temp)
	if actual != false {
		t.Errorf("TryDequeue failed to return false")
	}
}

func TestQueue_Length_Is_Initially_Zero(t *testing.T) {
	obj := NewQueue[int]()
	expected := 0
	if actual := obj.Length(); actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}

func TestQueue_Length_Is_One_After_Single_Enqueue(t *testing.T) {
	obj := NewQueue[int]()
	obj.Enqueue(1)
	expected := 1
	if actual := obj.Length(); actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}

func TestQueue_Length_Is_Correct_After_Multiple_Enqueues(t *testing.T) {
	obj := NewQueue[int]()
	expected := 50
	for i := 0; i < expected; i++ {
		obj.Enqueue(i)
	}
	if actual := obj.Length(); actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}

func TestQueue_Preserves_Ordering(t *testing.T) {
	obj := NewQueue[int]()

	expected1 := 10
	expected2 := 20
	expected3 := 30

	obj.Enqueue(expected1)
	obj.Enqueue(expected2)
	obj.Enqueue(expected3)

	actual1 := 0
	actual2 := 0
	actual3 := 0

	_ = obj.TryDequeue(&actual1)
	_ = obj.TryDequeue(&actual2)
	_ = obj.TryDequeue(&actual3)

	if actual1 != expected1 {
		t.Errorf("expected: %d, actual: %d", expected1, actual1)
	}
	if actual2 != expected2 {
		t.Errorf("expected: %d, actual: %d", expected2, actual2)
	}
	if actual3 != expected3 {
		t.Errorf("expected: %d, actual: %d", expected3, actual3)
	}
}

func BenchmarkQueue_Enqueue(b *testing.B) {

	obj := NewQueue[int]()

	for i := 0; i < b.N; i++ {
		obj.Enqueue(i)
	}
}

func BenchmarkQueue_Enqueue_TryDequeue(b *testing.B) {

	obj := NewQueue[int]()

	for i := 0; i < b.N; i++ {
		obj.Enqueue(i)
	}
	t := 0
	for i := 0; i < b.N; i++ {
		_ = obj.TryDequeue(&t)
	}
}
