package main

func mergeSort(list []int) []int {

	if len(list) <= 1 {
		return list
	} else if len(list) <= 2 {
		if list[0] > list[1] {
			list[0], list[1] = list[1], list[0]
		}

		return list
	} else {

		half := len(list) / 2

		// divide
		a := mergeSort(list[0:half])
		b := mergeSort(list[half:])

		buf := make([]int, len(list))

		i, j := 0, 0

		for i < len(a) && j < len(b) {
			if a[i] < b[j] {
				buf[i+j] = a[i]
				i += 1
			} else {
				buf[i+j] = b[j]
				j += 1
			}
		}

		for k, v := range a[i:] {
			buf[i+j+k] = v
		}

		for k, v := range b[j:] {
			buf[i+j+k] = v
		}

		// conquer
		return buf
	}
}
