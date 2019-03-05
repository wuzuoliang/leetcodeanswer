package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
*/

func Test_search(t *testing.T) {
	Convey("Test_search", t, func() {
		Convey("[4,5,6,7,0,1,2] 0", func() {
			So(IntShouldEqual(search([]int{4, 5, 6, 7, 0, 1, 2}, 0), 4), ShouldBeTrue)
		})

		Convey("[4,5,6,7,0,1,2] 1", func() {
			So(IntShouldEqual(search([]int{4, 5, 6, 7, 0, 1, 2}, 1), 5), ShouldBeTrue)
		})

		Convey("[4,5,6,7,0,1,2] 3", func() {
			So(IntShouldEqual(search([]int{4, 5, 6, 7, 0, 1, 2}, 3), -1), ShouldBeTrue)
		})

		Convey("[1,3] 3", func() {
			So(IntShouldEqual(search([]int{1, 3}, 3), 1), ShouldBeTrue)
		})
	})
}
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < nums[right] {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}
