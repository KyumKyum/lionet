// Stats and Logger types for database-wide observability.
package lionet

import "github.com/KyumKyum/lionet/v1/internal/cache"

type Stats struct {
	WALWrites      uint64
	MemtableSize   int64
	ImmutableCount int

	BlockCache cache.CacheStats

	LevelCount    int
	LevelSizes    []int64
	LevelTables   []int
	TotalSSTables int

	CompactionsRun     uint64
	CompactionsPending int
}

type Logger interface {
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type Compressor = interface {
	Compress(dst, src []byte) []byte
	Decompress(dst, src []byte) ([]byte, error)
	Name() string
}
