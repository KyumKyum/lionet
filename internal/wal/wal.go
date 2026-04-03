// Write-Ahead Log: persists mutations before Memtable apply for crash recovery.
package wal

import (
	"os"
	"sync"
)

type WAL struct {
	mu   sync.Mutex
	file *os.File
	dir  string
	seq  uint64
	sync bool
}

func Open(dir string, syncWrites bool) (*WAL, error) {
	panic("not implemented")
}

func (w *WAL) Append(entry *Entry) error {
	panic("not implemented")
}

func (w *WAL) Seq() uint64 {
	return w.seq
}

func (w *WAL) Close() error {
	panic("not implemented")
}

func (w *WAL) Rotate() error {
	panic("not implemented")
}

func (w *WAL) Truncate(upToSeq uint64) error {
	panic("not implemented")
}
