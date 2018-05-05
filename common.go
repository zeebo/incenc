package incenc

import (
	"errors"

	"github.com/zeebo/errs"
)

// Error wraps all of the returned errors.
var Error = errs.Class("incenc")

// findSuffix returns the length of the prefix of last and the suffix of value
// that is not present in last.
func findSuffix(last, value string) (int, string) {
	len_last, len_value := uint(len(last)), uint(len(value))

	// for loops are not inlinable, but goto's are. we also check that i >= 0
	// so that bounds checks are removed from the comparisons.
	i := uint(0)
loop:
	if i >= len_value {
		return int(i), ""
	}
	if i >= len_last || last[i] != value[i] {
		return int(i), value[i:]
	}
	i++
	goto loop
}

// bufferTooSmall is the error these functiosn returns to ensure that they can
// be inlined. A non-leaf function currently cannot be inlined, so we must
// return the same error every time.
var bufferTooSmall = errors.New("buffer too small")

// consume attempts to read n bytes from the buffer.
func consume(in []byte, n int) (out, data []byte, err error) {
	if n < 0 || len(in) < n {
		return nil, nil, bufferTooSmall
	}
	return in[n:], in[:n], nil
}

// readVarint reads an encoded unsigned integer from the byte buffer, and
// returns the integer and number of bytes read
func readVarint(in []byte) (out []byte, x uint16, err error) {
	if len(in) < 1 {
		return nil, 0, bufferTooSmall
	}
	b := in[0]
	if b < 128 {
		return in[1:], uint16(b), nil
	}
	if len(in) < 2 {
		return nil, 0, bufferTooSmall
	}
	return in[2:], uint16(b&127) | uint16(in[1])<<7, nil
}

// writeVarint writes an encoded unsigned integer to the byte buffer and
// returns a resliced buffer.
func writeVarint(in []byte, x uint16) (out []byte) {
	if x < 128 {
		return append(in, byte(x))
	} else {
		return append(in, 128|byte(x), byte(x>>7))
	}
}
