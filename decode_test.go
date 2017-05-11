package incenc

import "testing"

func TestDecode(t *testing.T) {
	data := []byte("\x00hello\x00\x05.world\x00")
	d := NewDecoder(data)

	d.Iterate(make([]byte, 100), func(name []byte) {
		t.Logf("%s", name)
	})
}
