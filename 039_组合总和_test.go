package Code

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/*
给你一个 无重复元素 的整数数组candidates 和一个目标整数target，找出candidates中可以使数字和为目标数target 的 所有不同组合 ，
并以列表形式返回。你可以按 任意顺序 返回这些组合。

candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

对于给定的输入，保证和为target 的不同组合数少于 150 个。

示例1：

输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。
示例2：

输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]]
示例 3：

输入: candidates = [2], target = 1
输出: []

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// https://www.bilibili.com/video/BV1KT4y1M7HJ
func Test_combinationSum(t *testing.T) {
	convey.Convey("Test_combinationSum", t, func() {

		convey.Convey("Input: candidates = [2,3,6,7], target = 7", func() {
			input := []int{2, 3, 6, 7}
			combinationSum(input, 7)

		})

		convey.Convey("Input: candidates = [2,3,5], target = 8", func() {
			input := []int{2, 3, 5}
			combinationSum(input, 8)
		})

		convey.Convey("Input: candidates = [8,7,4,3], target = 11", func() {
			input := []int{8, 7, 4, 3}
			fmt.Println(combinationSum(input, 11))
		})
	})
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	var dfs func(start int, temp []int, sum int)

	dfs = func(start int, temp []int, sum int) {
		if sum >= target {
			if sum == target {
				newTmp := make([]int, len(temp))
				copy(newTmp, temp)
				res = append(res, newTmp)
			}
			return
		}
		for i := start; i < len(candidates); i++ {
			temp = append(temp, candidates[i])
			//  关键点:不用i+1了，表示可以重复读取当前的数
			dfs(i, temp, sum+candidates[i])
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0, []int{}, 0)
	return res
}
