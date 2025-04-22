package threadsafe

import (
	"sync"
)

type Queue[T any] struct {
	lock sync.RWMutex
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		lock: sync.RWMutex{},
		data: make([]T, 0),
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.data = append(q.data, item)
}

func (q *Queue[T]) TryDequeue(item *T) bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.data == nil {
		q.data = make([]T, 0)
		return false
	}
	if len(q.data) == 0 {
		//fmt.Println("queue is empty")
		return false
	}

	temp := q.data[0]

	if len(q.data) == 1 {
		q.data = make([]T, 0)
	} else {
		q.data = q.data[1:]
	}

	*item = temp
	return true
}

func (q *Queue[_]) Length() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return len(q.data)

}
