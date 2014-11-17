package main

func merge(a, b []int) []int {
	buf := make([]int, len(a)+len(b))
	i, j, k := 0, 0, 0
	for ; i < len(a) && j < len(b); k++ {
		if a[i] < b[j] {
			buf[k] = a[i]
			i++
		} else {
			buf[k] = b[j]
			j++
		}
	}
	if i < len(a) {
		copy(buf[k:], a[i:])
	} else {
		copy(buf[k:], b[j:])
	}

	return buf
}
