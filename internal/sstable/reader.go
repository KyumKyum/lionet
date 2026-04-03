// SSTable reader: provides random and sequential read access to an SSTable file.
package sstable

import "os"

type Reader struct {
	file   *os.File
	index  *indexBlock
	filter *filterBlock
	meta   *TableMeta
}

func OpenReader(path string) (*Reader, error) {
	panic("not implemented")
}

func (r *Reader) Get(key []byte) ([]byte, bool, error) {
	// TODO:
	// 1. Check bloom filter
	// 2. Binary search index block
	// 3. Read & search data block
	panic("not implemented")
}

func (r *Reader) NewIterator() *TableIterator {
	panic("not implemented")
}

func (r *Reader) Close() error {
	panic("not implemented")
}
