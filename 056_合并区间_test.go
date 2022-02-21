package Code

import (
	"sort"
	"testing"
)

/**
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test56(t *testing.T) {
	t.Log(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	t.Log(merge([][]int{{1, 4}, {4, 5}}))
	t.Log(merge([][]int{{1, 4}, {2, 3}}))
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 第一个数字升序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)

	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if right >= intervals[i][0] {
			// 有交叉
			right = Max(right, intervals[i][1])
		} else {
			res = append(res, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		}
	}
	res = append(res, []int{left, right})
	return res
}
