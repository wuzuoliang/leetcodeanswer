package Code

import (
	"container/heap"
	"testing"
)

/**
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。



示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]


提示：

1 <= nums.length <= 105
k 的取值范围是 [1, 数组中不相同的元素的个数]
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/top-k-frequent-elements
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test347(t *testing.T) {
	t.Log(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))      // 1,2
	t.Log(topKFrequent([]int{1}, 1))                     // 1
	t.Log(topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2)) // -1 2
	t.Log(topKFrequent([]int{3, 0, 1, 0}, 1))            // 0
	t.Log(topKFrequent([]int{1, 1, 1, 2, 2, 3333}, 2))   // 1,2
}

func topKFrequentBest(nums []int, k int) []int {
	mp := map[int]int{}
	for _, i := range nums {
		mp[i]++
	}
	var maxTimes int
	for _, v := range mp {
		if v > maxTimes {
			maxTimes = v
		}
	}

	var res []int
	for k > 0 {
		for i, v := range mp {
			if v == maxTimes {
				res = append(res, i)
				k--
			}
		}
		maxTimes--
	}
	return res
}

func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
