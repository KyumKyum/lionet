// Iterator interface for scanning key/value pairs in sorted order.
package lionet

// Iterator iterates over key/value pairs in sorted order.
// Callers must check Valid() before accessing Key() or Value().
type Iterator interface {
	SeekToFirst()
	SeekToLast()
	Seek(target []byte)
	Next()
	Prev()
	Valid() bool
	Key() []byte
	Value() []byte
	Error() error
	Close() error
}

type IteratorOptions struct {
	LowerBound []byte
	UpperBound []byte
	Reverse    bool
}
