package main

import "fmt"

/**
首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
重复第二步，直到所有元素均排序完毕。

分析它的时间复杂度发现，无论最好最差的情况，其比较次数都是一样的多，第i趟排序需要进行n-i次关键字的比较。
而对于交换次数而言，当最好的时候，交换为0次，最差的时候，也就初始降序时，交换次数为n-1次，
基于最终的排序时间是比较与交换的次数总和，因此，总的时间复杂度依然为O(n^2)。虽然与冒泡排序同为O(n^2)，但选择排序的性能上还是要略优于冒泡排序的
*/
func main() {
	numbers := []uint64{5, 23, 1, 6, 7, 9, 2}
	ChooseSort(numbers)
	fmt.Println(numbers)
	fmt.Println("---")
	numbers2 := []uint64{5, 23, 1, 6, 7, 9, 2}
	ChooseSortV2(numbers2)
	fmt.Println(numbers2)
}

func ChooseSort(numbers uint64Slice) {
	for i := 0; i < len(numbers)-1; i++ {
		// 记录最小值位置
		min := i

		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[min] {
				min = j
			}
		}
		if i != min {
			numbers.swap(i, min)
		}
	}
}

// 上面的算法需要从某个数开始，一直扫描到尾部，我们可以优化算法，使得复杂度减少一半。
// 我们每一轮，除了找最小数之外，还找最大数，然后分别和前面和后面的元素交换，这样循环次数减少一半
func ChooseSortV2(numbers uint64Slice) {
	n := len(numbers)
	// 只需循环一半
	for i := 0; i < n/2; i++ {
		// 记录最小值、最大值位置
		minIndex := i
		maxIndex := i
		// 在这一轮迭代中要找到最大值和最小值的下标
		for j := i + 1; j < n-i; j++ {
			if numbers[j] > numbers[maxIndex] {
				maxIndex = j // 这一轮这个是大的，直接 continue
				continue
			}

			if numbers[j] < numbers[minIndex] {
				minIndex = j
			}
		}

		if maxIndex == i && minIndex != n-i-1 {
			// 如果最大值是开头的元素，而最小值不是最尾的元素
			// 先将最大值和最尾的元素交换
			numbers[n-i-1], numbers[maxIndex] = numbers[maxIndex], numbers[n-i-1]
			// 然后最小的元素放在最开头
			numbers[i], numbers[minIndex] = numbers[minIndex], numbers[i]
		} else if maxIndex == i && minIndex == n-i-1 {
			// 如果最大值在开头，最小值在结尾，直接交换
			numbers[minIndex], numbers[maxIndex] = numbers[maxIndex], numbers[minIndex]
		} else {
			// 否则先将最小值放在开头，再将最大值放在结尾
			numbers[i], numbers[minIndex] = numbers[minIndex], numbers[i]
			numbers[n-i-1], numbers[maxIndex] = numbers[maxIndex], numbers[n-i-1]
		}
	}
}

type uint64Slice []uint64

// 交换方法
func (numbers uint64Slice) swap(i, j int) {
	numbers[i], numbers[j] = numbers[j], numbers[i]
}
