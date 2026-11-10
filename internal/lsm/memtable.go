package lsm

import "sync"

type Memtable struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewMemtable() *Memtable {
	return &Memtable{
		data: make(map[string]string),
	}
}

func (m *Memtable) Set(key, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *Memtable) Get(key string) (string, bool) {
	m.mu.Lock()
	defer m.mu.RUnlock()
	val, exists := m.data[key]
	return val, exists
}

func (m *Memtable) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}
