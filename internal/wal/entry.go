// WAL entry encoding and on-disk format definitions.
package wal

import "encoding/binary"

type EntryType uint8

const (
	EntryPut    EntryType = 0x01
	EntryDelete EntryType = 0x02
	EntryBatch  EntryType = 0x03
)

// Entry on-disk format:
//
//	[checksum: 4B][length: 4B][seq: 8B][type: 1B][payload...]
//
// Put payload:    [key_len: 4B][key][value_len: 4B][value]
// Delete payload: [key_len: 4B][key]
type Entry struct {
	Seq   uint64
	Type  EntryType
	Key   []byte
	Value []byte
}

const (
	checksumSize = 4
	lengthSize   = 4
	seqSize      = 8
	typeSize     = 1
	headerSize   = checksumSize + lengthSize + seqSize + typeSize
)

func (e *Entry) EncodedSize() int {
	size := headerSize + 4 + len(e.Key)
	if e.Type == EntryPut {
		size += 4 + len(e.Value)
	}
	return size
}

func (e *Entry) Encode(dst []byte) {
	_ = dst
	_ = binary.LittleEndian
	// TODO: implement encoding with CRC32 checksum
	panic("not implemented")
}
