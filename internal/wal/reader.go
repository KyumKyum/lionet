// WAL replay reader used for crash recovery on startup.
package wal

type Reader struct {
	dir string
}

func NewReader(dir string) *Reader {
	return &Reader{dir: dir}
}

func (r *Reader) Replay(fn func(entry *Entry) error) error {
	panic("not implemented")
}
