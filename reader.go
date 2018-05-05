package incenc

// Reader decodes incremental encoded things.
type Reader struct {
	scratch []byte
}

// NewReaderWith constructs a reader using the provided scratch space.
func NewReaderWith(scratch []byte) Reader {
	return Reader{scratch: scratch[:0]}
}

// Reset clears the state of the Reader.
func (r *Reader) Reset() {
	r.scratch = r.scratch[:0]
}

// Next consumes the next value out of in, returns it as out, and the value as
// value. Value is only valid until the next call to Next.
func (r *Reader) Next(in []byte) (out, value []byte, err error) {
	in, prefix_len, err := readVarint(in)
	if err != nil {
		return nil, nil, err
	}

	in, value_len, err := readVarint(in)
	if err != nil {
		return nil, nil, err
	}

	in, suffix, err := consume(in, int(value_len))
	if err != nil {
		return nil, nil, err
	}

	// we read once to avoid having to worry about write barriers when updating
	// the slice.
	scratch := r.scratch

	_, scratch, err = consume(scratch, int(prefix_len))
	if err != nil {
		return nil, nil, err
	}
	scratch = append(scratch, suffix...)

	// save the appended value for the next call.
	r.scratch = scratch

	return in, scratch, nil
}
