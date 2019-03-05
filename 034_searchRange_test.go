package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
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
