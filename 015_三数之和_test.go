package Code

import (
	"sort"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

/**
给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
示例 2：

输入：nums = []
输出：[]
示例 3：

输入：nums = [0]
输出：[]


提示：

0 <= nums.length <= 3000
-105 <= nums[i] <= 105


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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
