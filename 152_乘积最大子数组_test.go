package Code

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
给你一个整数数组 nums，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

测试用例的答案是一个32-位 整数。

子数组 是数组的连续子序列。



示例 1:

输入: nums = [2,3,-2,4]
输出: 6
解释:子数组 [2,3] 有最大乘积 6。
示例 2:

输入: nums = [-2,0,-1]
输出: 0
解释:结果不能为 2, 因为 [-2,-1] 不是子数组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-product-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test152(t *testing.T) {
	convey.Convey("Test153", t, func() {
		convey.Convey("2,3,-2,4", func() {
			convey.So(maxProduct([]int{2, 3, -2, 4}), convey.ShouldEqual, 6)
		})
		convey.Convey("-2,0,-1", func() {
			convey.So(maxProduct([]int{-2, 0, -1}), convey.ShouldEqual, 0)
		})
		convey.Convey("5,2,-1,1,2,3,4", func() {
			convey.So(maxProduct([]int{5, 2, -1, 1, 2, 3, 4}), convey.ShouldEqual, 24)
		})
	})
}

/**
令imax为当前最大值，则当前最大值为 imax = max(imax * nums[i], nums[i])
由于存在负数，那么会导致最大的变最小的，最小的变最大的。因此还需要维护当前最小值imin，imin = min(imin * nums[i], nums[i])
当负数出现时则imax与imin进行交换再进行下一步计算

作者：guanpengchn
链接：https://leetcode-cn.com/problems/maximum-product-subarray/solution/hua-jie-suan-fa-152-cheng-ji-zui-da-zi-xu-lie-by-g/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = Max(mx*nums[i], Max(nums[i], mn*nums[i]))
		minF = Min(mn*nums[i], Min(nums[i], mx*nums[i]))
		ans = Max(maxF, ans)
	}
	return ans
}
