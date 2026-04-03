// Integration tests for Phase 1 completion criteria.
package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/KyumKyum/lionet/v1"
)

// TestMillionKeyWriteRead is the Phase 1 gate:
// Write 1M keys, close, reopen, read all 1M back with zero data loss.
func TestMillionKeyWriteRead(t *testing.T) {
	dir, err := os.MkdirTemp("", "lionet-test-million-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	const numKeys = 1_000_000

	db, err := lionet.Open(lionet.DefaultOptions(dir))
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}

	for i := 0; i < numKeys; i++ {
		key := []byte(fmt.Sprintf("key-%010d", i))
		val := []byte(fmt.Sprintf("value-%010d", i))
		if err := db.Put(key, val); err != nil {
			t.Fatalf("put failed at %d: %v", i, err)
		}
	}

	if err := db.Close(); err != nil {
		t.Fatalf("close failed: %v", err)
	}

	db, err = lionet.Open(lionet.DefaultOptions(dir))
	if err != nil {
		t.Fatalf("failed to reopen db: %v", err)
	}
	defer db.Close()

	for i := 0; i < numKeys; i++ {
		key := []byte(fmt.Sprintf("key-%010d", i))
		expected := []byte(fmt.Sprintf("value-%010d", i))

		val, err := db.Get(key)
		if err != nil {
			t.Fatalf("get failed at %d: %v", i, err)
		}
		if string(val) != string(expected) {
			t.Fatalf("mismatch at %d: got %q, want %q", i, val, expected)
		}
	}
}
