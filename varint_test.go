package incenc

import "testing"

func TestVarint(t *testing.T) {
	for i := 0; i < 32768; i++ {
		var buf [2]byte
		n := writeVarint(buf[:], uint16(i))
		if i < 128 && n == 2 {
			t.Errorf("%x", buf)
			t.Errorf("%v", buf)
			t.Fatal(i, "wrote too many bytes")
		}
		x, n := readVarint(buf[:])
		if i < 128 && n == 2 {
			t.Errorf("%x", buf)
			t.Errorf("%v", buf)
			t.Fatal(i, "read too many bytes")
		}
		if int(x) != i {
			t.Errorf("%x", buf)
			t.Errorf("%v", buf)
			t.Fatal(i, "failed to round trip", x)
		}
	}
}
