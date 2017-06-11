package incenc

// Decoder decodes incremental encoded things.
type Decoder struct {
	scratch []byte
}

// NewDecoder returns an incremental decoder.
func NewDecoder() *Decoder {
	return NewDecoderWith(nil)
}

// NewDecoderWith returns a new incremental decoder using the scratch buf.
func NewDecoderWith(scratch []byte) *Decoder {
	return &Decoder{
		scratch: scratch[:0],
	}
}

// Next consumes the next value out of in, returns it as out, and the value as
// value. Value is only valid until the next call to Next.
func (d *Decoder) Next(in []byte) (out, value []byte) {
	plen, amount := readVarint(in)
	in = in[amount:]

	d.scratch = d.scratch[:plen]

	for in[0] > 0 {
		d.scratch = append(d.scratch, in[0])
		in = in[1:]
	}
	in = in[1:]

	return in, d.scratch
}
