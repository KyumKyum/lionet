> [!NOTE]
> This is a project just kickstarted. Please wait a bit and we will provide an awesome DB! :)

# LionetDB

A pure Go embeddable key-value storage engine optimized for AI agent state — fast burst writes, prefix-scoped iteration, and automatic TTL-based expiry with zero manual GC. Classic LSM-tree architecture, zero cgo dependencies, `go get` and go.

> This is a preview. (A broad picture)

```bash
go get github.com/KyumKyum/lionet
```

```go
db, _ := lionetdb.Open("/tmp/mydb", lionetdb.DefaultOptions)
defer db.Close()

// Basic KV operations
db.Put([]byte("session:abc:state"), []byte(`{"step":3,"context":"..."}`))
val, _ := db.Get([]byte("session:abc:state"))

// Write with TTL — expired data is reclaimed automatically
db.Put([]byte("scratch:xyz"), []byte("temp data"), lionetdb.WithTTL(30 * time.Minute))

// Prefix scan — only reads relevant SSTables
iter := db.NewIterator(lionetdb.WithPrefix([]byte("session:abc:")))
defer iter.Close()
for iter.Seek(nil); iter.Valid(); iter.Next() {
    fmt.Printf("%s → %s\n", iter.Key(), iter.Value())
}
```

---

## Why LionetDB?

AI agents need local, persistent storage — session state, action logs, scratchpad memory, tool call history.
LionetDB uses a **classic LSM-tree** (like LevelDB), showing an advantage for sequential I/O for range scans.

**TBD**

---

## Features

**Core Engine**

- Classic LSM-tree with leveled compaction
- WAL for crash safety — `SIGKILL` recovery with zero data loss
- Bloom Filters for fast point lookups
- Prefix-compressed SSTables with configurable block size
- Pluggable compression (`Snappy`, `Zstd`, or bring your own)

**Agent-Optimized**

- **TTL-aware compaction** — expired SSTables are dropped without reading
- **Namespace TTL** — `db.SetNamespaceTTL("session:", 30*time.Minute)`
- **Prefix-scoped iteration** — skips irrelevant SSTables at every level
- **Recency-biased caching** — recent data (L0/L1) stays hot
- **Built-in observability** — compaction events, write stall notifications, per-namespace stats

**Coming Soon**

- Built-in Raft replication for distributed agent state
- Pluggable secondary index interface

---

## Configuration

```go
opts := lionetdb.Options{
    Dir:                  "/tmp/mydb",
    MemtableSize:         4 * 1024 * 1024,  // 4MB
    L0CompactionTrigger:  4,
    MaxLevels:            7,
    BloomBitsPerKey:      10,
    SyncWrites:           true,
    Compressor:           lionetdb.SnappyCompressor{},
}

db, err := lionetdb.Open(opts.Dir, opts)
```

For agent workloads, the defaults are tuned for: burst writes up to 10K ops/sec, sessions with 1KB–10KB state, and TTLs from minutes to hours.

---

## Project Status

LionetDB is under active development.

- [x] Project architecture and documentation
- [ ] Phase 1: Core LSM engine (WAL, Memtable, SSTable, Compaction)
- [ ] Phase 2: Agent extensions (TTL, Namespace, Hot/Cold separation)
- [ ] Phase 3: Pluggable interfaces (Compressor, SecondaryIndex, CompactionStrategy)
- [ ] Phase 4: Built-in Raft replication

---

## Contributing

We welcome contributions. See [CONTRIBUTING.md](./CONTRIBUTING.md) for workflow and guidelines.

---

## Contributors

| Name                   | Role                   |
| ---------------------- | ---------------------- |
| Lim Kyu Min (임규민)   | Maintainer, PL         |
| Han Seung Woo (한승우) | Maintainer             |
| Park Chul Wan (박철완) | Maintainer             |
| Lee Geon Hak (이건학)  | Maintainer             |
| Sung Hyeon Ju (성현주) | Maintainer             |

---

## License

Apache 2.0 — see [LICENSE](./LICENSE).
