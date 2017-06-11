package incenc

// Encoder encodes incremental encoding things.
type Encoder struct {
	last    string
	scratch [8]byte
}

// NewEncoder returns an Encoder for incremental encoding.
func NewEncoder() *Encoder {
	return &Encoder{}
}

// Append appends the value to the buf using the last value as state for
// reducing the amount of data needed to be written.
func (e *Encoder) Append(buf []byte, value string) []byte {
	// figure out how many bytes to use of last value
	i := commonPrefix(value, e.last)
	n := 1 + (i >> 7)

	// figure out size of varint
	start := len(buf)
	end := start + n + len(value) - i + 1

	// allocate more space if necessary
	if end >= cap(buf) {
		newb := make([]byte, len(buf), end+len(buf))
		copy(newb, buf)
		buf = newb[:end]
	} else {
		buf = buf[:end]
	}

	buf[start] = byte(i)
	if i > 127 {
		buf[start] |= 128
		buf[start+1] = byte(i >> 7)
	}
	copy(buf[start+n:], value[i:])
	buf[end-1] = 0

	e.last = value
	return buf
}
