package incenc

type Encoder struct {
	buf     []byte
	last    string
	scratch [8]byte
}

func NewEncoder() *Encoder {
	return NewEncoderWith(nil)
}

func NewEncoderWith(buf []byte) *Encoder {
	return &Encoder{
		buf: buf[:0],
	}
}

func (e *Encoder) Add(name string) {
	// figure out how many bytes to use of last value
	i := commonPrefix(name, e.last)
	n := 1 + (i >> 7)

	// figure out size of varint
	start := len(e.buf)
	end := start + n + len(name) - i + 1

	// allocate more space if necessary
	if end >= cap(e.buf) {
		newb := make([]byte, len(e.buf), end+len(e.buf))
		copy(newb, e.buf)
		e.buf = newb[:end]
	} else {
		e.buf = e.buf[:end]
	}

	e.buf[start] = byte(i)
	if i > 127 {
		e.buf[start] |= 128
		e.buf[start+1] = byte(i >> 7)
	}
	copy(e.buf[start+n:], name[i:])
	e.buf[end-1] = 0

	e.last = name
}

func (e *Encoder) Bytes() []byte {
	return e.buf
}
