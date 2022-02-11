package Code

import (
	"sort"
	"testing"
)

/**
给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。

注意:

可以认为区间的终点总是大于它的起点。
区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。
示例 1:

输入: [ [1,2], [2,3], [3,4], [1,3] ]
输出: 1
解释: 移除 [1,3] 后，剩下的区间没有重叠。

示例 2:

输入: [ [1,2], [1,2], [1,2] ]
输出: 2
解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。

示例 3:

输入: [ [1,2], [2,3] ]
输出: 0
解释: 你不需要移除任何区间，因为它们已经是无重叠的了。

*/

func Test435(t *testing.T) {
	t.Log(intervalSchedule([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	t.Log(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
}

// 贪心算法，区间调度
func intervalSchedule(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] < intervals[j][1] {
			return true
		} else {
			if intervals[i][0] <= intervals[j][0] {
				return true
			}
		}
		return false
	})

	// 肯定会有一个在里面
	count := 1
	curEnd := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= curEnd {
			count++
			curEnd = intervals[i][1]
		}
	}
	return count
}

// 无重叠区间
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	return n - intervalSchedule(intervals)
}
