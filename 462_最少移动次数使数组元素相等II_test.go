package Code

import (
	"sort"
	"testing"
)

/**
给你一个长度为 n 的整数数组 nums ，返回使所有数组元素相等需要的最少移动数。

在一步操作中，你可以使数组中的一个元素加 1 或者减 1 。



示例 1：

输入：nums = [1,2,3]
输出：2
解释：
只需要两步操作（每步操作指南使一个元素加 1 或减 1）：
[1,2,3]  =>  [2,2,3]  =>  [2,2,2]
示例 2：

输入：nums = [1,10,2,9]
输出：16


提示：

n == nums.length
1 <= nums.length <= 105
-109 <= nums[i] <= 109

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test462(t *testing.T) {
	t.Log(minMoves2([]int{1, 2, 3, 4, 5, 6}))
}

// 找中位数
func minMoves2(nums []int) (ans int) {
	sort.Ints(nums)
	x := nums[len(nums)/2]
	for _, num := range nums {
		ans += abs(num - x)
	}
	return
}
