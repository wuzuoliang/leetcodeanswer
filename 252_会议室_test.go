package Code

import (
	"sort"
	"testing"
)

/**
给你输入若干形如[begin, end]的区间，代表若干会议的开始时间和结束时间，请你计算至少需要申请多少间会议室。
如果把每个会议的起止时间看做一个线段区间，那么题目就是让你求最多有几个重叠区间
*/
func Test252(t *testing.T) {
	t.Log(minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}))
}

func minMeetingRooms(meetings [][]int) int {
	if len(meetings) == 0 {
		return 0
	}
	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i][1] < meetings[j][1] {
			return true
		} else {
			if meetings[i][0] <= meetings[j][0] {
				return true
			}
		}
		return false
	})
	countMax := 1
	for i := 1; i < len(meetings)-1; i++ {
		count := 1
		for j := 0; j < i; j++ {
			if meetings[j][0] < meetings[i][1] {
				count++
			}
		}
		if count > countMax {
			countMax = count
		}
	}
	return countMax
}
