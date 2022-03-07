package Code

import (
	"strconv"
	"testing"
)

/**
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。



示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]


提示：

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test46(t *testing.T) {
	t.Log(permute([]int{1, 2, 3}))
}

var res [][]int
var innerres []int

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	innerres = make([]int, 0)
	visited := make(map[string]struct{})
	used := make(map[int]struct{})
	dfs46(nums, used, visited, len(nums))
	return res
}

func dfs46(nums []int, used map[int]struct{}, visited map[string]struct{}, n int) {
	if len(innerres) == n {
		str := ""
		newValues := make([]int, 0, n)
		for _, v := range innerres {
			str += strconv.Itoa(v)
			newValues = append(newValues, v)
		}
		if _, ok := visited[str]; ok {
			return
		}
		res = append(res, newValues)
		return
	}

	for _, v := range nums {
		if _, ok := used[v]; ok {
			continue
		}

		used[v] = struct{}{}
		innerres = append(innerres, v)

		dfs46(nums, used, visited, n)

		innerres = (innerres)[0 : len(innerres)-1]
		delete(used, v)
	}
}
