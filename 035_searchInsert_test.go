package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
Given a sorted array and a target value, return the index if the target is found.
If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:

Input: [1,3,5,6], 5
Output: 2
Example 2:

Input: [1,3,5,6], 2
Output: 1
Example 3:

Input: [1,3,5,6], 7
Output: 4
Example 4:

Input: [1,3,5,6], 0
Output: 0
*/

func Test_searchInsert(t *testing.T) {
	Convey("Test_searchInsert", t, func() {
		Convey("Input: [1,3,5,6], 5", func() {
			So(IntShouldEqual(searchInsert([]int{1, 3, 5, 6}, 5), 2), ShouldBeTrue)
		})
		Convey("Input: [1,3,5,6], 2", func() {
			So(IntShouldEqual(searchInsert([]int{1, 3, 5, 6}, 2), 1), ShouldBeTrue)
		})
		Convey("Input: [1,3,5,6], 7", func() {
			So(IntShouldEqual(searchInsert([]int{1, 3, 5, 6}, 7), 4), ShouldBeTrue)
		})
		Convey("Input: [1,3,5,6], 0", func() {
			So(IntShouldEqual(searchInsert([]int{1, 3, 5, 6}, 0), 0), ShouldBeTrue)
		})
	})
}

func searchInsert(nums []int, target int) int {
	var ret int
	for i := range nums {
		if nums[i] >= target {
			ret = i
			break
		}
	}
	if len(nums) > 0 && nums[0] < target && ret == 0 {
		ret = len(nums)
	}
	return ret
}
