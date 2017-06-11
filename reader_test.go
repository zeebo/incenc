package incenc

import "testing"

func TestReader(t *testing.T) {
	data := []byte("\x00\x05hello\x05\x06.world")
	var value []byte
	var r Reader

	for len(data) > 0 {
		data, value = r.Next(data)
		t.Logf("%s: %x", value, value)
	}
}

func BenchmarkReader(b *testing.B) {
	scratch := make([]byte, 256)
	buf := encodeCorpus(nil)
	b.SetBytes(int64(len(buf)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		in := buf
		r := Reader{Scratch: scratch}

		for len(in) > 0 {
			in, _ = r.Next(in)
		}
	}
}
