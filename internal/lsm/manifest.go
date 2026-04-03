// Manifest: append-only log tracking SSTable layout across levels (version edits).
package lsm

import (
	"os"
	"sync"
)

type Manifest struct {
	mu   sync.Mutex
	file *os.File
	dir  string
}

type VersionEdit struct {
	AddedTables   []TableRef
	RemovedTables []TableRef
	NextFileNum   uint64
	LastSeq       uint64
}

type TableRef struct {
	Level   int
	FileNum uint64
}

func OpenManifest(dir string) (*Manifest, error) {
	panic("not implemented")
}

func (m *Manifest) LogEdit(edit *VersionEdit) error {
	panic("not implemented")
}

func (m *Manifest) Recover() (*VersionSet, error) {
	panic("not implemented")
}

func (m *Manifest) Close() error {
	panic("not implemented")
}
