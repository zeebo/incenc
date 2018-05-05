// +build gofuzz

package incenc

func Fuzz(data []byte) int {
	var r Reader
	var err error

	for len(data) > 0 {
		data, _, err = r.Next(data)
		if err != nil {
			return 0
		}
	}

	return 1
}
