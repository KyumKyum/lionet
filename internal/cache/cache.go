// Cache: block cache interface for reducing SSTable I/O.
package cache

type Cache interface {
	Get(key uint64) []byte
	Set(key uint64, value []byte)
	Delete(key uint64)
	Close()
	Stats() CacheStats
}

type CacheStats struct {
	Hits      uint64
	Misses    uint64
	Evictions uint64
	TotalSize int64
}
