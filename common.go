package incenc

// commonPrefix returns how many bytes of the two input strings are shared.
func commonPrefix(left, right string) (i int) {
	for i = 0; i < len(left) && i < len(right); i++ {
		if left[i] != right[i] {
			return i
		}
	}
	return i
}
