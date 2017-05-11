package incenc

// readVarint reads an encoded signed integer from the byte buffer, and returns
// the integer and number of bytes read
func readVarint(r []byte) (x uint16, n int) {
	b := r[0]
	if b < 128 {
		return uint16(b), 1
	}
	return uint16(b&127) | uint16(r[1])<<7, 2
}

func writeVarint(r []byte, x uint16) int {
	r[0] = byte(x)
	if x < 128 {
		return 1
	}
	r[0] |= 128
	r[1] = byte(x >> 7)
	return 2
}
