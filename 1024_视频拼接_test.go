package Code

import (
	"sort"
	"testing"
)

func Test1024(t *testing.T) {
	t.Log(videoStitching([][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}, 10))
}

// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247491102&idx=1&sn=755633b7d225970135cd6c8b2f500a27&scene=21#wechat_redirect
func videoStitching(clips [][]int, T int) int {
	if len(clips) == 0 {
		return 0
	}
	// 按起点升序排列，起点相同的降序排列
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
	// 这段代码的时间复杂度是多少呢？虽然代码中有一个嵌套的 while 循环，但这个嵌套 while 循环的时间复杂度是O(N)。因为当i递增到n时循环就会结束，所以这段代码只会执行O(N)次。
	// 但是别忘了我们对clips数组进行了一次排序，消耗了O(NlogN)的时间，所以本算法的总时间复杂度是O(NlogN)。
}
