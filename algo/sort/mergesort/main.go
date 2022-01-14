package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
算法步骤
申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列；
设定两个指针，最初位置分别为两个已经排序序列的起始位置；
比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置；
重复步骤 3 直到某一指针达到序列尾；
将另一序列剩下的所有元素直接复制到合并序列尾。
https://mmbiz.qpic.cn/mmbiz_gif/k5430ljpYPMpkFUr3BCNATOIkthtqYDYGGo93Rwe4tyLZpb9xq8AFI1E7xHB7C4YzjeYItpSwKeXjM6s5DN99g/640?wx_fmt=gif&tp=webp&wxfrom=5&wx_lazy=1

将两个有序数组进行合并，最多进行 n 次比较就可以生成一个新的有序数组，n 是两个数组长度较大的那个。

归并操作最坏的时间复杂度为：O(n)，其中 n 是较长数组的长度。

归并操作最好的时间复杂度为：O(n)，其中 n 是较短数组的长度。

每次都是一分为二，特别均匀，所以最差和最坏时间复杂度都一样。归并操作的时间复杂度为：O(n)，因此总的时间复杂度为：T(n)=2T(n/2)+O(n)，根据主定理公式可以知道时间复杂度为：O(nlogn)。
*/
func main() {
	list := InitSlice()
	fmt.Println(list)
	MergeSort1(list, 0, len(list))
	fmt.Println(list)

	fmt.Println("---")

	list = InitSlice()
	fmt.Println(list)
	MergeSort2(list, 0, len(list))
	fmt.Println(list)

	fmt.Println("---")

	list = InitSlice()
	fmt.Println(list)
	MergeSort3(list, len(list))
	fmt.Println(list)

}

func init() {
	rand.Seed(time.Now().Unix())
}

func InitSlice() []int {
	list := make([]int, 10)
	for i := 0; i < 10; i++ {
		list[i] = rand.Intn(101)
	}
	return list
}

// 归并操作
func merge(array []int, begin int, mid int, end int) {
	// 申请额外的空间来合并两个有序数组，这两个数组是 array[begin,mid),array[mid,end)
	leftSize := mid - begin         // 左边数组的长度
	rightSize := end - mid          // 右边数组的长度
	newSize := leftSize + rightSize // 辅助数组的长度
	result := make([]int, 0, newSize)

	l, r := 0, 0
	for l < leftSize && r < rightSize {
		lValue := array[begin+l] // 左边数组的元素
		rValue := array[mid+r]   // 右边数组的元素
		// 小的元素先放进辅助数组里
		if lValue < rValue {
			result = append(result, lValue)
			l++
		} else {
			result = append(result, rValue)
			r++
		}
	}

	// 将剩下的元素追加到辅助数组后面
	result = append(result, array[begin+l:mid]...)
	result = append(result, array[mid+r:end]...)

	// 将辅助数组的元素复制回原数组，这样该辅助空间就可以被释放掉
	for i := 0; i < newSize; i++ {
		array[begin+i] = result[i]
	}
	return
}

// 自顶向下归并排序，排序范围在 [begin,end) 的数组
func MergeSort1(array []int, begin int, end int) {
	// 元素数量大于1时才进入递归
	if end-begin > 1 {

		// 将数组一分为二，分为 array[begin,mid) 和 array[mid,high)
		mid := begin + (end-begin+1)/2

		// 先将左边排序好
		MergeSort1(array, begin, mid)

		// 再将右边排序好
		MergeSort1(array, mid, end)

		// 两个有序数组进行合并
		merge(array, begin, mid, end)
	}
}

// 自底向上归并排序, 非递归排序，没有递归那样程序栈的增加，效率比自顶向下的递归版本高
func MergeSort2(array []int, begin, end int) {
	// 步数为1开始，step长度的数组表示一个有序的数组
	step := 1

	// 范围大于 step 的数组才可以进入归并
	for end-begin > step {
		// 从头到尾对数组进行归并操作
		// step << 1 = 2 * step 表示偏移到后两个有序数组将它们进行归并
		for i := begin; i < end; i += step << 1 {
			var lo = i                // 第一个有序数组的上界
			var mid = lo + step       // 第一个有序数组的下界，第二个有序数组的上界
			var hi = lo + (step << 1) // 第二个有序数组的下界

			// 不存在第二个数组，直接返回
			if mid > end {
				return
			}

			// 第二个数组长度不够
			if hi > end {
				hi = end
			}

			// 两个有序数组进行合并
			merge(array, lo, mid, hi)
		}

		// 上面的 step 长度的两个数组都归并成一个数组了，现在步长翻倍
		step <<= 1
	}
}

func InsertSort(list []int) {
	n := len(list)
	// 进行 N-1 轮迭代
	for i := 1; i <= n-1; i++ {
		deal := list[i] // 待排序的数
		j := i - 1      // 待排序的数左边的第一个数的位置

		// 如果第一次比较，比左边的已排好序的第一个数小，那么进入处理
		if deal < list[j] {
			// 一直往左边找，比待排序大的数都往后挪，腾空位给待排序插入
			for ; j >= 0 && deal < list[j]; j-- {
				list[j+1] = list[j] // 某数后移，给待排序留空位
			}
			list[j+1] = deal // 结束了，待排序的数插入空位
		}
	}
}

// 自底向上归并排序优化版本,因为手摇只多了逆序翻转的操作，时间复杂度是 O(n)，虽然时间复杂度稍稍多了一点，但存储空间复杂度降为了 O(1)
func MergeSort3(array []int, n int) {
	// 按照三个元素为一组进行小数组排序，使用直接插入排序
	blockSize := 3
	a, b := 0, blockSize
	for b <= n {
		InsertSort(array[a:b])
		a = b
		b += blockSize
	}
	InsertSort(array[a:n])

	// 将这些小数组进行归并
	for blockSize < n {
		a, b = 0, 2*blockSize
		for b <= n {
			merge2(array, a, a+blockSize, b)
			a = b
			b += 2 * blockSize
		}
		if m := a + blockSize; m < n {
			merge2(array, a, m, n)
		}
		blockSize *= 2
	}
}

// 原地归并操作
func merge2(array []int, begin, mid, end int) {
	// 三个下标，将数组 array[begin,mid] 和 array[mid,end-1]进行原地归并
	i, j, k := begin, mid, end-1 // 因为数组下标从0开始，所以 k = end-1

	for j-i > 0 && k-j >= 0 {
		step := 0
		// 从 i 向右移动，找到第一个 array[i]>array[j]的索引
		for j-i > 0 && array[i] <= array[j] {
			i++
		}

		// 从 j 向右移动，找到第一个 array[j]>array[i]的索引
		for k-j >= 0 && array[j] <= array[i] {
			j++
			step++
		}

		// 进行手摇翻转，将 array[i,mid] 和 [mid,j-1] 进行位置互换
		// mid 是从 j 开始向右出发的，所以 mid = j-step
		rotation(array, i, j-step, j-1)
		i = i + step
	}

}

// 手摇算法，将 array[l,l+1,l+2,...,mid-2,mid-1,mid,mid+1,mid+2,...,r-2,r-1,r] 从mid开始两边交换位置
// 1.先逆序前部分：array[mid-1,mid-2,...,l+2,l+1,l]
// 2.后逆序后部分：array[r,r-1,r-2,...,mid+2,mid+1,mid]
// 3.上两步完成后：array[mid-1,mid-2,...,l+2,l+1,l,r,r-1,r-2,...,mid+2,mid+1,mid]
// 4.整体逆序： array[mid,mid+1,mid+2,...,r-2,r-1,r,l,l+1,l+2,...,mid-2,mid-1]
func rotation(array []int, l, mid, r int) {
	reverse(array, l, mid-1)
	reverse(array, mid, r)
	reverse(array, l, r)
}

func reverse(array []int, l, r int) {
	for l < r {
		// 左右互相交换
		array[l], array[r] = array[r], array[l]
		l++
		r--
	}
}
