// LRU: LRU eviction policy block cache (Phase 1 default).
package cache

// LRUCache implements Cache using a least-recently-used eviction policy.
type LRUCache struct {
	// TODO: implement LRU cache
}

func NewLRUCache(maxBytes int64) *LRUCache {
	panic("not implemented")
}

func (c *LRUCache) Get(key uint64) []byte              { panic("not implemented") }
func (c *LRUCache) Set(key uint64, value []byte)       { panic("not implemented") }
func (c *LRUCache) Delete(key uint64)                  { panic("not implemented") }
func (c *LRUCache) Close()                             { panic("not implemented") }
func (c *LRUCache) Stats() CacheStats                  { panic("not implemented") }
