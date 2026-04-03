# ADR - 001: MVP 

## Status: ACCEPTED

## Context
- The phases and mvp of current project (Good start will result a good project)
- Draw a big picture, but Finish one by one.

## Decision

### Phase 1: Core LSM Engine
- WAL, Memtable, SSTable, LSM-tree manager, R/W path, pulbic API 
- Completion: 1M key write/read with zero data loss and SIGKILL crash recovery.

### Phase 2: Agent-optimized extensions
- TTL-aware compaction, namespace TTL policies, prefix-optimized iteration, hot/cold separation with a recency biased block cache. 


