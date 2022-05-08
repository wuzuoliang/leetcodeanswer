package Code

import (
	"testing"
)

/**
给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。

数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。



示例 1：

输入：nums = [4,6,7,7]
输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
示例 2：

输入：nums = [4,4,3,2,1]
输出：[[4,4]]


提示：

1 <= nums.length <= 15
-100 <= nums[i] <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/increasing-subsequences
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test491(t *testing.T) {
	t.Log(findSubsequences([]int{4, 6, 7, 7}))
	t.Log(findSubsequences([]int{4, 4, 3, 2, 1}))
	t.Log(findSubsequences([]int{4, 4, 3, 4, 2, 1}))
}

func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	var dfs func(res *[][]int, nums []int, start int, subStr []int)
	dfs = func(res *[][]int, nums []int, start int, subStr []int) {
		if len(subStr) > 1 {
			tmp := make([]int, len(subStr))
			copy(tmp, subStr)
			*res = append(*res, tmp)
		}
		history := [201]int{} // 这里使用数组来进行去重操作，题目说数值范围[-100, 100]
		for index := start; index < len(nums); index++ {
			if len(subStr) != 0 {
				if nums[index] < subStr[len(subStr)-1] {
					continue
				}
			}
			if history[nums[index]+100] == 1 {
				continue
			}
			history[nums[index]+100] = 1 // 记录这个元素在本层用过了，本层后面不能再用了
			dfs(res, nums, index+1, append(subStr, nums[index]))
		}
	}
	dfs(&res, nums, 0, nil)
	return res
}
