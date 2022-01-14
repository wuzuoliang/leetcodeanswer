package main

// ConcurrentQuickSort ...
func ConcurrentQuickSort(a *[]int, left int, right int, chanSend chan int) {
	/**
	// 改进版
	if (right - left) < 1000 {
		QuickSort(a, left, right)
		chanSend <- 0
		return
	}
	*/
	if right <= left {
		chanSend <- 0
		return
	}

	idx := left

	// Pick a pivot
	pivotIndex := left

	// Move the pivot to the right
	swap(a, pivotIndex, right)

	// Pile elements smaller than the pivot on the left
	for i := left; i < right; i++ {
		if (*a)[i] < (*a)[right] {
			swap(a, i, idx)
			idx++
		}
	}

	// Place the pivot after the last smaller element
	swap(a, idx, right)

	chanReceive := make(chan int)

	// Go down the rabbit hole
	go ConcurrentQuickSort(a, left, idx, chanReceive)
	go ConcurrentQuickSort(a, idx+1, right, chanReceive)
	<-chanReceive
	<-chanReceive
	chanSend <- 0
	return
}
