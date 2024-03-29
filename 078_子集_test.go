package Code

import (
	"testing"
)

/**
给你一个整数数组nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。



示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
示例 2：

输入：nums = [0]
输出：[[],[0]]


提示：

1 <= nums.length <= 10
-10 <= nums[i] <= 10
nums 中的所有元素 互不相同

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test78(t *testing.T) {
	t.Log(subsets([]int{1, 2, 3}))
}

var subss [][]int
var subs []int

func subsets(nums []int) [][]int {
	backtrace78(nums, 0)
	return subss
}

// https://labuladong.gitee.io/algo/4/30/112/
func backtrace78(nums []int, start int) {
	newSub := make([]int, len(subs))
	copy(newSub, subs)
	subss = append(subss, newSub)

	// 注意 i 从 start 开始递增
	for i := start; i < len(nums); i++ {
		// 做选择
		subs = append(subs, nums[i])
		// 回溯
		backtrace78(nums, i+1)
		// 撤销选择
		subs = subs[:len(subs)-1]
	}
}
