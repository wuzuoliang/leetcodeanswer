package Code

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为[4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回-1。



示例 1：

输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
示例2：

输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
示例 3：

输入：nums = [1], target = 0
输出：-1


提示：

1 <= nums.length <= 5000
-10^4 <= nums[i] <= 10^4
nums 中的每个值都 独一无二
题目数据保证 nums 在预先未知的某个下标上进行了旋转
-10^4 <= target <= 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test_search(t *testing.T) {
	convey.Convey("Test_search", t, func() {
		convey.Convey("[4,5,6,7,0,1,2] 0", func() {
			convey.So(IntShouldEqual(search([]int{4, 5, 6, 7, 0, 1, 2}, 0), 4), convey.ShouldBeTrue)
		})

		convey.Convey("[4,5,6,7,0,1,2] 1", func() {
			convey.So(IntShouldEqual(search([]int{4, 5, 6, 7, 0, 1, 2}, 1), 5), convey.ShouldBeTrue)
		})

		convey.Convey("[4,5,6,7,0,1,2] 3", func() {
			convey.So(IntShouldEqual(search([]int{4, 5, 6, 7, 0, 1, 2}, 3), -1), convey.ShouldBeTrue)
		})

		convey.Convey("[1,3] 3", func() {
			convey.So(IntShouldEqual(search([]int{1, 3}, 3), 1), convey.ShouldBeTrue)
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
