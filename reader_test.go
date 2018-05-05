package incenc

import "testing"

func TestReader(t *testing.T) {
	data := []byte("\x00\x05hello\x05\x06.world")
	var value []byte
	var r Reader
	var err error

	for len(data) > 0 {
		data, value, err = r.Next(data)
		assertNoError(t, err)
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
		r := NewReaderWith(scratch)

		for len(in) > 0 {
			in, _, _ = r.Next(in)
		}
	}
}
