package threadsafe

import "sync"

type Flag struct {
	mutex *sync.RWMutex
	value bool
}

func NewFlag() *Flag {
	return &Flag{
		mutex: &sync.RWMutex{},
		value: false,
	}
}

func (f *Flag) Clear() {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.value = false
}

func (f *Flag) Set() {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.value = true
}

func (f *Flag) IsSet() bool {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return f.value
}
