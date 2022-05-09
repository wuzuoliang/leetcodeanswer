package Code

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

/**
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。



示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组[4,-1,2,1] 的和最大，为6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [5,4,-1,7,8]
输出：23


提示：

1 <= nums.length <= 105
-104 <= nums[i] <= 104


进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test_maxSubArray(t *testing.T) {
	convey.Convey("Test_maxSubArray", t, func() {
		convey.So(IntShouldEqual(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), 6), convey.ShouldBeTrue)

		convey.So(IntShouldEqual(maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), 6), convey.ShouldBeTrue)
	})
}

// 复杂度分析：时间复杂度：O(N)。空间复杂度：O(N)
func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	result := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = Max(dp[i-1]+nums[i], nums[i])
		result = Max(dp[i], result)
	}
	return result
}

// 贪心算法 时间复杂度：O(n)  空间复杂度：O(1)
func maxSubArray2(nums []int) int {
	res := 0
	tmpSum := 0
	for i := 0; i < len(nums); i++ {
		tmpSum += nums[i]
		// 取区间累计的最大值（相当于不断确定最大子序终止位置）
		if tmpSum > res {
			res = tmpSum
		} else if tmpSum <= 0 {
			// 相当于重置最大子序起始位置，因为遇到负数一定是拉低总和
			tmpSum = 0
		}
	}
	return res
}
