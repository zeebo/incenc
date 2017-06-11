package incenc

import "testing"

func TestDecode(t *testing.T) {
	data := []byte("\x00hello\x00\x05.world\x00")
	var value []byte

	d := NewDecoder()

	for len(data) > 0 {
		data, value = d.Next(data)
		t.Logf("%s", value)
	}
}

func BenchmarkDecode(b *testing.B) {
	scratch := make([]byte, 256)
	buf := encodeCorpus(nil)
	b.SetBytes(int64(len(buf)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		in := buf
		d := NewDecoderWith(scratch)
		for len(in) > 0 {
			in, _ = d.Next(in)
		}
	}
}
