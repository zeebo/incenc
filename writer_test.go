package incenc

import (
	"bytes"
	"testing"
)

func TestWriter(t *testing.T) {
	var buf []byte
	var w Writer

	buf = w.Append(buf, "hello")
	buf = w.Append(buf, "hello.world")

	exp := []byte("\x00\x05hello\x05\x06.world")

	if !bytes.Equal(exp, buf) {
		t.Errorf("exp: %x", exp)
		t.Errorf("got: %x", buf)
	}
}

func encodeCorpus(buf []byte) []byte {
	var w Writer
	for _, v := range corpus {
		buf = w.Append(buf, v)
	}
	return buf
}

func BenchmarkWriter(b *testing.B) {
	buf := make([]byte, 1<<20)
	b.SetBytes(int64(corpusLength))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf = buf[:0]
		encodeCorpus(buf)
	}
}
