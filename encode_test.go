package incenc

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	var buf []byte
	e := NewEncoder()

	buf = e.Append(buf, "hello")
	buf = e.Append(buf, "hello.world")

	exp := []byte("\x00hello\x00\x05.world\x00")

	if !bytes.Equal(exp, buf) {
		t.Errorf("exp: %x", exp)
		t.Errorf("got: %x", buf)
	}
}

func encodeCorpus(buf []byte) []byte {
	e := NewEncoder()
	for _, v := range corpus {
		buf = e.Append(buf, v)
	}
	return buf
}

func BenchmarkEncode(b *testing.B) {
	buf := make([]byte, 1<<20)
	b.SetBytes(int64(corpusLength))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf = buf[:0]
		encodeCorpus(buf)
	}
}
