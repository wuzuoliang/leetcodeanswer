package Code

import (
	"math/rand"
	"testing"
	"time"
)

/**
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。



示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test215(t *testing.T) {
	//t.Log(findKthLargest([]int{3,2,1,5,6,4},2))//5
	//t.Log(findKthLargest([]int{3,2,3,1,2,4,5,5,6},4))//4
	t.Log(findKthLargest2([]int{5, 2, 4, 1, 3, 6, 0}, 4)) //3
}

// 大顶堆
func findKthLargest2(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}

// 快排
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSort(nums, 0, len(nums)-1, k)
}

func quickSort(nums []int, low, high, k int) int {
	pivot := randomPartition(nums, low, high)
	if pivot == k-1 {
		return nums[pivot]
	}
	if pivot > k-1 {
		return quickSort(nums, low, pivot-1, k)
	}
	return quickSort(nums, pivot+1, high, k)
}

func randomPartition(nums []int, low, high int) int {
	i := rand.Int()%(high-low+1) + low
	nums[i], nums[low] = nums[low], nums[i]
	return partition(nums, low, high)
}

func partition(nums []int, low, high int) int {
	pivot := nums[low]

	for low < high {
		for low < high && nums[high] <= pivot {
			high--
		}
		nums[low] = nums[high]

		for low < high && nums[low] >= pivot {
			low++
		}
		nums[high] = nums[low]
	}

	nums[low] = pivot

	return low
}
