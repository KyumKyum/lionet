// Snapshot provides a consistent read-only view of the DB at a point in time.
package lionet

type Snapshot struct {
	db      *DB
	version uint64
}

func (s *Snapshot) Get(key []byte) ([]byte, error) {
	panic("not implemented")
}

func (s *Snapshot) Has(key []byte) (bool, error) {
	panic("not implemented")
}

func (s *Snapshot) NewIterator(opts *IteratorOptions) Iterator {
	panic("not implemented")
}

func (s *Snapshot) Release() {
	panic("not implemented")
}
