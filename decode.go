package incenc

type Decoder struct {
	data []byte
}

func NewDecoder(data []byte) *Decoder {
	return &Decoder{
		data: data,
	}
}

func (d *Decoder) Iterate(buf []byte, cb func([]byte)) {
	if buf != nil {
		buf = buf[:0]
	}

	for len(d.data) > 0 {
		plen, amount := readVarint(d.data)
		d.data = d.data[amount:]

		buf = buf[:plen]
		for len(d.data) > 0 {
			b := d.data[0]
			d.data = d.data[1:]

			if b == 0 {
				break
			}

			buf = append(buf, b)
		}

		cb(buf)
	}
}
