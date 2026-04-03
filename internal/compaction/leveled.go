// Leveled: leveled compaction strategy (default), targeting size ratio between levels.
package compaction

// LeveledCompaction implements Strategy using size-tiered leveled compaction.
type LeveledCompaction struct {
	// TODO: implement leveled compaction
}

func (l *LeveledCompaction) PickCompaction(levels []LevelMeta) *CompactionTask {
	panic("not implemented")
}

func (l *LeveledCompaction) Name() string {
	return "leveled"
}
