package incenc

// Writer encodes incremental encoding things.
type Writer struct {
	last string
}

// Append appends the value to the buf using the last value as state for
// reducing the amount of data needed to be written.
func (w *Writer) Append(buf []byte, value string) []byte {
	// figure out how many bytes to use of last value
	prefix_len := commonPrefix(value, w.last)
	value_len := len(value) - prefix_len

	// compute sizes of varints
	prefix_varint := 1 + (prefix_len >> 7)
	value_varint := 1 + (len(value) >> 7)

	// figure out the range of the buffer we'll be using
	start := len(buf)
	end := start + prefix_varint + value_varint + value_len

	// allocate more space if necessary
	if end >= cap(buf) {
		newb := make([]byte, len(buf), end+len(buf))
		copy(newb, buf)
		buf = newb[:end]
	} else {
		buf = buf[:end]
	}

	writeVarint(buf[start:], uint16(prefix_len))
	writeVarint(buf[start+prefix_varint:], uint16(value_len))
	copy(buf[start+prefix_varint+value_varint:], value[prefix_len:])

	w.last = value
	return buf
}

// AppendBytes is like Append except the value is a []byte instead of a string.
func (w *Writer) AppendBytes(buf []byte, value []byte) []byte {
	// TODO(jeff): man go has some problems with []byte and string. jeeze.
	// we can avoid this allocation at the cost of a bunch of duplication,
	// or we can avoid the duplication at the cost of this allocation. i don't
	// know if we can do both.
	return w.Append(buf, string(value))
}
