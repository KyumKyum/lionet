// Compressor: pluggable block compression interface with a noop default.
package codec

type Compressor interface {
	Compress(dst, src []byte) []byte
	Decompress(dst, src []byte) ([]byte, error)
	Name() string
}

type NoopCompressor struct{}

func (NoopCompressor) Compress(dst, src []byte) []byte              { return src }
func (NoopCompressor) Decompress(dst, src []byte) ([]byte, error)   { return src, nil }
func (NoopCompressor) Name() string                                  { return "noop" }
