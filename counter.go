package threadsafe

import "sync"

type Counter struct {
	lock  *sync.Mutex
	value uint64
}

func NewCounter() *Counter {
	return &Counter{
		lock:  &sync.Mutex{},
		value: 0,
	}
}

func (cc *Counter) Increment() uint64 {
	cc.lock.Lock()
	defer cc.lock.Unlock()
	cc.value++
	return cc.value
}
