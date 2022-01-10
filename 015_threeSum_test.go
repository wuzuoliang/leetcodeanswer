package Code

import (
	"sort"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

/**
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
https://www.geekxh.com/1.0.数组系列/008.html#_02、题目分析
*/

func Test_threeSum(t *testing.T) {
	Convey("Test_threeSum", t, func() {

		Convey("[-1, 0, 1, 2, -1, -4]", func() {
			t.Log(threeSum([]int{-1, 0, 1, 2, -1, -4}))
		})

	})

}

func threeSum(nums []int) [][]int {
	var results [][]int
	length := len(nums)
	if length <= 3 {
		return results
	}
	sort.Ints(nums)
	for start := 0; start+2 < length; start++ {
		if start > 1 && nums[start] == nums[start-1] {
			start++
			continue
		}
		left := start + 1
		right := length - 1
		if nums[start] > 0 {
			break
		}
		for left < right {
			if nums[start]+nums[left]+nums[right] == 0 {
				results = append(results, []int{nums[start], nums[left], nums[right]})
				left++
			} else if nums[start]+nums[left]+nums[right] > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return results
}
