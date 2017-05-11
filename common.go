package incenc

func commonPrefix(left, right string) (i int) {
	for ; i < len(left) && i < len(right); i++ {
		if left[i] != right[i] {
			return i
		}
	}
	return i
}
