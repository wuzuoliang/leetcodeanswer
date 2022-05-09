package Code

import (
	"math"
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

// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247491640&idx=1&sn=60344b057f19a4765a15ed17cb7c8018&scene=21#wechat_redirect
func minMeetingRooms(meetings [][]int) int {
	n := len(meetings)
	if n == 0 {
		return 0
	}
	begins := make([]int, 0, n)
	ends := make([]int, 0, n)
	for _, v := range meetings {
		begins = append(begins, v[0])
		ends = append(ends, v[1])
	}
	sort.Ints(begins)
	sort.Ints(ends)

	max := math.MinInt
	count := 0
	l := 0
	r := 0
	for l < n && r < n {
		if begins[l] < ends[r] {
			count++
			l++
		} else {
			count--
			r++
		}
		max = Max(max, count)
	}
	return max
}
