// Memtable: in-memory sorted key/value store backed by a skiplist.
package memtable

import "sync"

type Memtable struct {
	mu       sync.RWMutex
	list     *Skiplist
	size     int64
	maxSize  int64
	readOnly bool
}

func New(maxSize int64) *Memtable {
	return &Memtable{
		list:    NewSkiplist(),
		maxSize: maxSize,
	}
}

func (m *Memtable) Put(key, value []byte) error {
	panic("not implemented")
}

func (m *Memtable) Delete(key []byte) error {
	panic("not implemented")
}

func (m *Memtable) Get(key []byte) ([]byte, bool) {
	panic("not implemented")
}

func (m *Memtable) ApproximateSize() int64 {
	return m.size
}

func (m *Memtable) IsFull() bool {
	return m.size >= m.maxSize
}

func (m *Memtable) MarkImmutable() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.readOnly = true
}

func (m *Memtable) NewIterator() *MemtableIterator {
	panic("not implemented")
}
