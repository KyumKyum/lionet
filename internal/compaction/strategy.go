// Strategy: pluggable compaction interface for selecting which SSTables to merge.
package compaction

import "github.com/KyumKyum/lionet/v1/internal/sstable"

type Strategy interface {
	PickCompaction(levels []LevelMeta) *CompactionTask
	Name() string
}

type LevelMeta struct {
	Level     int
	Tables    []sstable.TableMeta
	TotalSize uint64
}

type CompactionTask struct {
	InputLevel  int
	OutputLevel int
	Inputs      []sstable.TableMeta
	Overlaps    []sstable.TableMeta
}
