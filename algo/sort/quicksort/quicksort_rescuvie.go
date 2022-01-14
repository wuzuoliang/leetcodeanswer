package main

// QuickSort ...
func QuickSort(numbers *[]int, left int, right int) {
	if right <= left {
		return
	}

	idx := left

	// Pick a pivot
	pivotIndex := left

	// Move the pivot to the right
	swap(numbers, pivotIndex, right)

	// Pile elements smaller than the pivot on the left
	for i := left; i < right; i++ {
		if (*numbers)[i] < (*numbers)[right] {
			swap(numbers, i, idx)
			idx++
		}
	}

	// Place the pivot after the last smaller element
	swap(numbers, idx, right)
	QuickSort(numbers, left, idx)
	QuickSort(numbers, idx+1, right)
	return

}

func QuickSort2(numbers *[]int, left int, right int) {
	var middle int
	tempStart := left
	tempEnd := right

	if tempStart >= tempEnd {
		return
	}
	pivot := (*numbers)[left]
	for left < right {
		for left < right && (*numbers)[right] > pivot {
			right--
		}
		if left < right {
			swap(numbers, left, right)
			left++
		}
		for left < right && (*numbers)[left] < pivot {
			left++
		}
		if left < right {
			swap(numbers, left, right)
			right--
		}
	}
	(*numbers)[left] = pivot
	middle = left

	QuickSort(numbers, tempStart, middle-1)
	QuickSort(numbers, middle+1, tempEnd)
}
