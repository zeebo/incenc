package incenc

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	e := NewEncoder()

	e.Add("hello")
	e.Add("hello.world")

	got := e.Bytes()
	exp := []byte("\x00hello\x00\x05.world\x00")

	if !bytes.Equal(exp, got) {
		t.Errorf("exp: %x", exp)
		t.Errorf("got: %x", got)
	}
}

func BenchmarkEncode(b *testing.B) {
	buf := make([]byte, 1<<20)
	total := int64(0)
	for _, v := range corpus {
		total += int64(len(v))
	}
	b.SetBytes(total)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		e := NewEncoderWith(buf)
		for _, v := range corpus {
			e.Add(v)
		}
	}
}

var corpus = []string{
	"foo",
}
