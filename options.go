// Options for configuring the database at open time.
package lionet

import "github.com/KyumKyum/lionet/v1/internal/compaction"

type Options struct {
	Dir                   string
	MemtableSize          int64
	MaxLevels             int
	BlockSize             int
	BloomFilterBitsPerKey int
	CompactionStrategy    compaction.Strategy
	Compressor            Compressor
	BlockCacheSize        int64
	SyncWrites            bool
	MaxWriteBatchSize     int
	Logger                Logger
}

func DefaultOptions(dir string) Options {
	return Options{
		Dir:                   dir,
		MemtableSize:          64 << 20,
		MaxLevels:             7,
		BlockSize:             4 << 10,
		BloomFilterBitsPerKey: 10,
		BlockCacheSize:        256 << 20,
		SyncWrites:            true,
		MaxWriteBatchSize:     1000,
	}
}
