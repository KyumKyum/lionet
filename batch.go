// WriteBatch accumulates mutations to be applied atomically via DB.Write.
package lionet

type WriteBatch struct {
	ops []batchOp
}

type batchOpType uint8

const (
	batchOpPut batchOpType = iota
	batchOpDelete
)

type batchOp struct {
	opType batchOpType
	key    []byte
	value  []byte
}

func NewWriteBatch() *WriteBatch {
	return &WriteBatch{}
}

func (b *WriteBatch) Put(key, value []byte) {
	b.ops = append(b.ops, batchOp{opType: batchOpPut, key: key, value: value})
}

func (b *WriteBatch) Delete(key []byte) {
	b.ops = append(b.ops, batchOp{opType: batchOpDelete, key: key})
}

func (b *WriteBatch) Count() int {
	return len(b.ops)
}

func (b *WriteBatch) Reset() {
	b.ops = b.ops[:0]
}
