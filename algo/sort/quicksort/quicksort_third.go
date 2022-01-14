package main

/**
排序数列可能存在大量重复值，使用三向切分快速排序，将数组分成三部分，大于基准数，等于基准数，小于基准数，这个时候需要维护三个下标。

三切分，把小于基准数的扔到左边，大于基准数的扔到右边，相同的元素会进行聚集。

如果存在大量重复元素，排序速度将极大提高，将会是线性时间，因为相同的元素将会聚集在中间，这些元素不再进入下一个递归迭代
*/

// QuickSort4 三切分的快速排序
func QuickSort4(array []int, begin, end int) {
	if begin < end {
		// 三向切分函数，返回左边和右边下标
		lt, gt := partition3(array, begin, end)
		// 从lt到gt的部分是三切分的中间数列
		// 左边三向快排
		QuickSort4(array, begin, lt-1)
		// 右边三向快排
		QuickSort4(array, gt+1, end)
	}
}

// 切分函数，并返回切分元素的下标
func partition3(array []int, begin, end int) (int, int) {
	lt := begin       // 左下标从第一位开始
	gt := end         // 右下标是数组的最后一位
	i := begin + 1    // 中间下标，从第二位开始
	v := array[begin] // 基准数

	// 以中间坐标为准
	for i <= gt {
		if array[i] > v { // 大于基准数，那么交换，右指针左移
			array[i], array[gt] = array[gt], array[i]
			gt--
		} else if array[i] < v { // 小于基准数，那么交换，左指针右移
			array[i], array[lt] = array[lt], array[i]
			lt++
			i++
		} else {
			i++
		}
	}

	return lt, gt
}
