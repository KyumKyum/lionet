# Contributing to LionetDB

Thank you for your interest in contributing to LionetDB. This document explains how we work together.

## Branching Model (Gitflow)

We use **Gitflow**. Here's how branches work:

```
main          ← production-ready, tagged releases only
  │
  └─ develop  ← integration branch, all feature branches merge here
       │
       ├─ feat/wal-record-format      ← new features
       ├─ fix/bloom-false-positive     ← bug fixes
       ├─ test/crash-recovery-matrix   ← test additions
       ├─ docs/sstable-format-spec     ← documentation
       ├─ refactor/memtable-iterator   ← code improvements
       └─ bench/compaction-throughput  ← benchmarks
```

### Branch Rules

**`main`** — Always stable. Never push directly. Only merge from `develop` via a release PR. Every merge to `main` gets a version tag.

**`develop`** — Integration branch. All feature branches merge here. Should always build and pass tests, but may contain work-in-progress features.

**Feature branches** — Created from `develop`, merged back into `develop`. Use the naming convention below. Keep them short-lived — merge often, don't let branches live for weeks.

### Branch Naming

```
type/short-description

type:
  feat/     ← new feature
  fix/      ← bug fix
  test/     ← adding or improving tests
  docs/     ← documentation only
  refactor/ ← code change that doesn't add features or fix bugs
  bench/    ← benchmarks and performance work
```

Examples:
```
feat/wal-block-spanning-records
fix/sstable-index-off-by-one
test/concurrent-memtable-stress
docs/compaction-strategy-adr
refactor/skiplist-arena-allocator
bench/1m-key-write-throughput
```

### Release Flow

When `develop` is ready for release:

1. Create a release branch: `release/v0.1.0`
2. Only bug fixes go into the release branch (no new features).
3. Merge release branch into `main` and tag it: `v0.1.0`
4. Merge release branch back into `develop`.
5. Delete the release branch.

### Hotfix Flow

For critical bugs in `main`:

1. Create a hotfix branch from `main`: `hotfix/wal-corruption-on-restart`
2. Fix the bug, add a test for it.
3. Merge into `main` and tag: `v0.1.1`
4. Merge into `develop`.
5. Delete the hotfix branch.

---

## Making a Pull Request

### Before Opening a PR

- [ ] Your branch is up-to-date with `develop` (`git rebase develop` or `git merge develop`)
- [ ] All tests pass: `go test -race ./...`
- [ ] No linting errors: `go vet ./...`
- [ ] New code has tests (unit tests at minimum, crash tests for WAL/Flush/Compaction)
- [ ] Exported functions have godoc comments

### PR Guidelines

TBD

---

## Code Review

Every PR needs **at least one approval** before merging into `develop`.

**For reviewers:**
- Review the code, not the person. Be specific and kind.
- If you approve, it means: "I read this, I understand it, and I believe it is correct."
- If you disagree, explain your reasoning. If the thread gets long, move it to a call.
- Pay extra attention to: crash safety, concurrency, and data corruption risks. A small bug in a storage engine can mean silent data loss.

**For authors:**
- Respond to every comment, even if just with "done."
- Don't force-push after review has started (it makes it hard to see what changed).
- Squash commits when merging if the branch history is messy.

---

## Commit Messages

```
type(scope): short summary

type:  feat, fix, test, docs, refactor, bench
scope: wal, memtable, sstable, lsm, bloom, api, manifest (optional)
summary: present tense, lowercase, no period at end
```

Examples:
```
feat(wal): add CRC32 checksum to record format
fix(lsm): prevent tombstone drop before bottom level
test(sstable): add concurrent read/write stress test
docs(compaction): explain leveled compaction triggers
refactor(memtable): switch to arena allocator for skip list nodes
bench: add 1M key write throughput benchmark
```

Keep the summary line under 72 characters. If you need more detail, add a body after a blank line:

```
fix(lsm): prevent tombstone drop before bottom level

Tombstones were being dropped during L1→L2 compaction, causing
deleted keys to reappear from L3. Tombstones must only be dropped
at the bottom level where no older data can exist below.

Closes #42
```

---

## Testing Requirements

We take testing seriously. A storage engine bug can mean silent data loss.

### What Every PR Needs

| Change Type | Required Tests |
|---|---|
| New feature | Unit tests + integration test |
| Bug fix | Regression test that fails without the fix |
| WAL changes | Crash test (SIGKILL recovery) |
| Compaction changes | Crash test + correctness test (no data loss after compaction) |
| Concurrency changes | Stress test with `-race` flag |
| File format changes | Fuzz test for encoding/decoding roundtrip |

### Running Tests

```bash
# All tests
go test ./...

# With race detector (required before PR)
go test -race ./...

# Specific package
go test ./wal/...
go test ./sstable/...

# Benchmarks
go test -bench=. ./...

# Fuzz tests (run for 30 seconds)
go test -fuzz=FuzzWALRecord -fuzztime=30s ./wal/

# Verbose (see individual test names)
go test -v -race ./...
```

---

## Writing Rules

All technical terminology must remain in English. Do not transliterate into Korean.

- ✅ "Memtable이 가득 차면 Flush가 발생한다"
- ❌ "맴테이블이 가득 차면 플러시가 발생한다"
- ✅ "Footer에는 Index Block의 Offset이 저장된다"
- ❌ "풋터에는 인덱스 블록의 오프셋이 저장된다"

This applies to: code, comments, PR descriptions, commit messages, and documentation.

Use English in code, comments, and PR descriptions. Team chat can be in Korean.

---

## Architecture Decision Records (ADR)

When you make an important technical choice, record it in `docs/decisions/`:

```
docs/decisions/
├── 001-skiplist-over-rbtree.md
├── 002-classic-lsm-over-wisckey.md
├── 003-leveled-over-size-tiered.md
└── ...
```

Each ADR answers:
1. **What** did we decide?
2. **Why** did we choose this over alternatives?
3. **What** are the tradeoffs?

If you are about to make a choice that affects other modules, write the ADR first and get feedback before coding.

---

## Communication

| Purpose | Channel |
|---------|---------|
| Quick questions | Team chat |
| Design discussions | GitHub Discussions |
| Bug reports, feature requests | GitHub Issues |
| Architecture decisions | ADR + GitHub Discussion |
| Code review | GitHub PR comments |

---

## Getting Help

- Open a GitHub Discussion for general questions.
- Open a GitHub Issue for bugs or feature ideas.
- Ask in team chat for quick help.

Every question you ask helps us find gaps in our documentation. Asking is contributing.
