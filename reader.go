package incenc

// Reader decodes incremental encoded things.
type Reader struct {
	Scratch []byte
}

// Next consumes the next value out of in, returns it as out, and the value as
// value. Value is only valid until the next call to Next.
func (r *Reader) Next(in []byte) (out, value []byte) {
	prefix_len, size_1 := readVarint(in)
	value_len, size_2 := readVarint(in[size_1:])

	start := size_1 + size_2
	end := start + int(value_len)

	r.Scratch = r.Scratch[:prefix_len]
	r.Scratch = append(r.Scratch, in[start:end]...)

	return in[end:], r.Scratch
}
