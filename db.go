// DB is the main public API entry point. Safe for concurrent use.
package lionet

import "sync"

type DB struct {
	opts   Options
	mu     sync.RWMutex
	closed bool

	// TODO: internal components
	// wal       *wal.WAL
	// mem       *memtable.Memtable
	// imm       []*memtable.Memtable
	// lsm       *lsm.Manager
	// compactor *compaction.Compactor
	// cache     cache.Cache
}

func Open(opts Options) (*DB, error) {
	// TODO:
	// 1. Create/validate directory
	// 2. Load or create MANIFEST
	// 3. Open WAL, replay entries into Memtable
	// 4. Open existing SSTables per MANIFEST
	// 5. Start compaction goroutine
	// 6. Initialize block cache
	panic("not implemented")
}

func (db *DB) Close() error {
	// TODO:
	// 1. Signal compaction to stop
	// 2. Flush active memtable
	// 3. Close WAL
	// 4. Close all SSTable readers
	// 5. Sync MANIFEST
	panic("not implemented")
}

func (db *DB) Put(key, value []byte) error {
	panic("not implemented")
}

func (db *DB) Get(key []byte) ([]byte, error) {
	panic("not implemented")
}

func (db *DB) Delete(key []byte) error {
	panic("not implemented")
}

func (db *DB) Has(key []byte) (bool, error) {
	panic("not implemented")
}

func (db *DB) Write(batch *WriteBatch) error {
	panic("not implemented")
}

func (db *DB) NewIterator(opts *IteratorOptions) Iterator {
	panic("not implemented")
}

func (db *DB) NewSnapshot() *Snapshot {
	panic("not implemented")
}

func (db *DB) CompactRange(start, end []byte) error {
	panic("not implemented")
}

func (db *DB) Stats() *Stats {
	panic("not implemented")
}
