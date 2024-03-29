package Code

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

/**
题目189: 旋转数组
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
示例 1:

输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
1
2
解释:

向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入: [-1,-100,3,99] 和 k = 2
输出: [3,99,-1,-100]
1
2
解释:

向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]
说明:

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。

要求使用空间复杂度为 O(1) 的 原地 算法。
*/
func TestRotateArray(t *testing.T) {
	convey.Convey("TestRotateArray", t, func() {

		convey.Convey("1", func() {
			convey.So(IntSliceShouldEqual(rotateArray1([]int{1, 2, 3, 4, 5, 6, 7}, 3), []int{5, 6, 7, 1, 2, 3, 4}), convey.ShouldBeTrue)
		})

		convey.Convey("2", func() {
			convey.So(IntSliceShouldEqual(rotateArray1([]int{4, 5, 5, 9}, 2), []int{5, 9, 4, 5}), convey.ShouldBeTrue)
		})
	})
}

func rotateArray1(nums []int, k int) []int {
	reverse189(nums)
	reverse189(nums[:k%len(nums)])
	reverse189(nums[k%len(nums):])
	return nums
}

func reverse189(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}
