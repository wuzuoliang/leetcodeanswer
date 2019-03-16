package Code

func QuickSort(in []int) []int {
	quickSort(in, 0, len(in)-1)
	return in
}

func quickSort(in []int, l, r int) {
	if l >= r {
		return
	}
	pivot := in[l]
	left := l + 1
	right := r
	for l <= r {
		for left <= right && in[left] < pivot {

			left++
		}

		for left <= right && in[right] >= pivot {
			right--
		}
		if left > right {
			break
		}
		in[left], in[right] = in[right], in[left]

	}
	/* swap the smaller with pivot */
	in[l], in[right] = in[right], in[l]

	quickSort(in, l, right-1)
	quickSort(in, right+1, r)
}
