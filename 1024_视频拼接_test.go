package Code

import (
	"sort"
	"testing"
)

func Test1024(t *testing.T) {
	t.Log(videoStitching([][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}, 10))
}

func videoStitching(clips [][]int, T int) int {
	if len(clips) == 0 {
		return 0
	}
	sort.Slice(clips, func(i, j int) bool {
		if clips[i][0] < clips[j][0] {
			return true
		} else {
			if clips[i][1] <= clips[j][1] {
				return true
			}
		}
		return false
	})
	// 记录选择的短视频个数
	res := 0
	curEnd := 0
	nextEnd := 0
	i := 0
	n := len(clips)
	for i < n && clips[i][0] <= curEnd {
		// 在第 res 个视频的区间内贪心选择下一个视频
		for i < n && clips[i][0] <= curEnd {
			nextEnd = Max(nextEnd, clips[i][1])
			i++
		}
		// 找到下一个视频，更新 curEnd
		res++
		curEnd = nextEnd
		if curEnd >= T {
			// 已经可以拼出区间 [0, T]
			return res
		}
	}
	// 无法连续拼出区间 [0, T]
	return -1
}
