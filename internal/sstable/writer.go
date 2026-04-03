// SSTable writer: builds immutable sorted SSTable files from sorted key/value pairs.
// Layout: [data blocks...][filter block][index block][footer]
package sstable

import "os"

type Writer struct {
	file      *os.File
	block     *blockWriter
	index     *indexWriter
	filter    *filterWriter
	offset    uint64
	count     uint64
	blockSize int
}

func NewWriter(path string, blockSize int, bitsPerKey int) (*Writer, error) {
	panic("not implemented")
}

func (w *Writer) Add(key, value []byte) error {
	panic("not implemented")
}

func (w *Writer) Finish() (*TableMeta, error) {
	panic("not implemented")
}

func (w *Writer) Abort() error {
	panic("not implemented")
}

type TableMeta struct {
	Path        string
	FileSize    uint64
	KeyCount    uint64
	SmallestKey []byte
	LargestKey  []byte
}
