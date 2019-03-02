package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

func Test_TwoSum(t *testing.T) {
	Convey("Test_TwoSum", t, func() {
		Convey("Given nums = [2, 7, 11, 15], target = 9", func() {
			So(IntSliceShouldEqual(twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1}), ShouldBeTrue)
		})
	})
}
func twoSum(nums []int, target int) []int {
	n := len(nums)

	if n == 0 {
		return []int{}
	}

	m := make(map[int]int)

	for i := 0; i < n; i++ {
		m[nums[i]] = i
	}

	for i := 0; i < n; i++ {
		ans := target - nums[i]
		v, ok := m[ans]
		if ok && i != v {
			return []int{i, v}
		}
	}
	return []int{}
}
