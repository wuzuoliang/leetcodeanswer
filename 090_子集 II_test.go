package Code

import (
	"sort"
	"testing"
)

/**
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。



示例 1：

输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
示例 2：

输入：nums = [0]
输出：[[],[0]]


提示：

1 <= nums.length <= 10
-10 <= nums[i] <= 10


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test90(t *testing.T) {
	t.Log(subsetsWithDup([]int{1, 2, 2}))
}

var res90 [][]int

func subsetsWithDup(nums []int) [][]int {
	res90 = make([][]int, 0)
	sort.Ints(nums)
	dfs([]int{}, nums, 0)
	return res90
}
func dfs(temp, num []int, start int) {
	tmp := make([]int, len(temp))
	copy(tmp, temp)
	res90 = append(res90, tmp)

	for i := start; i < len(num); i++ {
		if i > start && num[i] == num[i-1] {
			continue
		}
		dfs(append(temp, num[i]), num, i+1)
	}
}
