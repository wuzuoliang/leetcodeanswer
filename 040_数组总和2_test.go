package Code

import (
	"sort"
	"testing"
)

/**
给定一个候选人编号的集合candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。

candidates中的每个数字在每个组合中只能使用一次。

注意：解集不能包含重复的组合。



示例1:

输入: candidates =[10,1,2,7,6,1,5], target =8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
示例2:

输入: candidates =[2,5,2,1,2], target =5,
输出:
[
[1,2,2],
[5]
]


提示:

1 <=candidates.length <= 100
1 <=candidates[i] <= 50
1 <= target <= 30

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test_combinationSum2(t *testing.T) {
	t.Log(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	var trcak []int
	var res [][]int
	sort.Ints(candidates)
	backtracking(0, 0, target, candidates, trcak, &res)
	return res
}
func backtracking(startIndex, sum, target int, candidates, trcak []int, res *[][]int) {
	//终止条件
	if sum == target {
		tmp := make([]int, len(trcak))
		//拷贝
		copy(tmp, trcak)
		//放入结果集
		*res = append(*res, tmp)
		return
	}
	//回溯
	for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ {
		// 若当前树层有使用过相同的元素，则跳过
		if i > startIndex && candidates[i] == candidates[i-1] {
			continue
		}
		//更新路径集合和sum
		trcak = append(trcak, candidates[i])
		backtracking(i+1, sum+candidates[i], target, candidates, trcak, res)
		//回溯
		trcak = trcak[:len(trcak)-1]
	}
}
