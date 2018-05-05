package incenc

// Writer encodes incremental encoding things.
type Writer struct {
	last string
}

// Reset clears the state of the Writer.
func (w *Writer) Reset() {
	w.last = ""
}

// Copy returns a safe copy of the current Writer with state preserved.
func (w *Writer) Copy() *Writer {
	return &Writer{last: w.last}
}

// Append appends the value to the buf using the last value as state for
// reducing the amount of data needed to be written.
func (w *Writer) Append(buf []byte, value string) ([]byte, error) {
	// figure out how many bytes to use of last value
	prefix_len, suffix := findSuffix(w.last, value)
	suffix_len := len(suffix)

	if prefix_len > 1<<15 || suffix_len > 1<<15 || prefix_len < 0 {
		return nil, Error.New("value too large: %q", value)
	}

	buf = writeVarint(buf, uint16(prefix_len))
	buf = writeVarint(buf, uint16(suffix_len))
	buf = append(buf, suffix...)

	w.last = value
	return buf, nil
}

// AppendBytes is like Append except the value is a []byte instead of a string.
func (w *Writer) AppendBytes(buf []byte, value []byte) ([]byte, error) {
	// TODO(jeff): man go has some problems with []byte and string. jeeze.
	// we can avoid this allocation at the cost of a bunch of duplication,
	// or we can avoid the duplication at the cost of this allocation. i don't
	// know if we can do both.
	return w.Append(buf, string(value))
}
