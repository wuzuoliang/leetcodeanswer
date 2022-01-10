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
*/

func Test_threeSum(t *testing.T) {
	Convey("Test_threeSum", t, func() {
		Convey("[-1, 0, 1, 2, -1, -4]", func() {
			So(IntSliceShouldEqual(threeSum([]int{-1, 0, 1, 2, -1, -4})[0], []int{-1, -1, 2}), ShouldBeTrue)
			So(IntSliceShouldEqual(threeSum([]int{-1, 0, 1, 2, -1, -4})[1], []int{-1, 0, 1}), ShouldBeTrue)
			So(IntShouldEqual(len(threeSum([]int{-1, 0, 1, 2, -1, -4})), 2), ShouldBeTrue)
		})
		Convey("[-2, 0, 2, 2, -2, 0]", func() {
			So(IntSliceShouldEqual(threeSum([]int{-2, 0, 2, 2, -2, 0})[0], []int{-2, 0, 2}), ShouldBeTrue)
			So(IntShouldEqual(len(threeSum([]int{-2, 0, 2, 2, -2, 0})), 1), ShouldBeTrue)
		})
		Convey("[0, 0, 0]", func() {
			So(IntSliceShouldEqual(threeSum([]int{0, 0, 0})[0], []int{0, 0, 0}), ShouldBeTrue)
		})
	})
}

func threeSum(nums []int) [][]int {
	var results [][]int
	length := len(nums)
	if length < 3 {
		return results
	}
	sort.Ints(nums)
	for start := 0; start+2 < length; start++ {
		left := start + 1
		right := length - 1
		if nums[start] > 0 {
			break
		}
		if start == 0 || nums[start] != nums[start-1] {
			for left < right {
				if nums[start]+nums[left]+nums[right] == 0 {
					results = append(results, []int{nums[start], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if nums[start]+nums[left]+nums[right] > 0 {
					right--
				} else {
					left++
				}
			}
		}
	}
	return results
}
