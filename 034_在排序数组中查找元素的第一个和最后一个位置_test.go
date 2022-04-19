package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

/**
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回[-1, -1]。

进阶：

你可以设计并实现时间复杂度为O(log n)的算法解决此问题吗？


示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test_searchRange(t *testing.T) {
	Convey("Test_searchRange", t, func() {
		Convey("[5,7,7,8,8,10] 8", func() {
			So(IntSliceShouldEqual(searchRange([]int{5, 7, 7, 8, 8, 10}, 8), []int{3, 4}), ShouldBeTrue)
		})

		Convey("[5,7,7,8,8,10] 6", func() {
			So(IntSliceShouldEqual(searchRange([]int{5, 7, 7, 8, 8, 10}, 6), []int{-1, -1}), ShouldBeTrue)
		})

		Convey("[] 6", func() {
			So(IntSliceShouldEqual(searchRange([]int{}, 6), []int{-1, -1}), ShouldBeTrue)
		})

		Convey("[1] 1", func() {
			So(IntSliceShouldEqual(searchRange([]int{1}, 1), []int{0, 0}), ShouldBeTrue)
		})
	})
}

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) < 1 {
		return res
	}
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[right] != target {
		return res
	}
	res[0] = right
	right = len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	res[1] = left - 1
	return res
}

func searchRangeLib(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}
