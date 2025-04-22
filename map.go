package threadsafe

import "sync"

type Map[K comparable, V any] struct {
	lock sync.RWMutex
	data map[K]V
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		data: make(map[K]V),
	}
}

func (m *Map[K, V]) Put(key K, value V) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.data[key] = value
}

func (m *Map[K, V]) Get(key K) (V, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	v, ok := m.data[key]
	return v, ok
}

func (m *Map[K, V]) GetAndDelete(key K) (V, bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	v, ok := m.data[key]
	if ok {
		delete(m.data, key)
	}
	return v, ok
}

func (m *Map[K, V]) Delete(keys ...K) {
	m.lock.Lock()
	defer m.lock.Unlock()
	for _, key := range keys {
		delete(m.data, key)
	}
}

func (m *Map[K, V]) Size() int {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return len(m.data)
}

func (m *Map[K, V]) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.data = make(map[K]V)
}

func (m *Map[K, V]) Exists(key K) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	_, ok := m.data[key]
	return ok
}

func (m *Map[K, V]) Iterate(f func(key K, value V) bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	for k, v := range m.data {
		if !f(k, v) {
			break
		}
	}
}

func (m *Map[K, V]) Keys() []K {
	m.lock.RLock()
	defer m.lock.RUnlock()
	keys := make([]K, 0)
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}

func (m *Map[K, V]) Values() []V {
	m.lock.RLock()
	defer m.lock.RUnlock()
	values := make([]V, 0)
	for _, v := range m.data {
		values = append(values, v)
	}
	return values
}
