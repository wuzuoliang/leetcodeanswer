package Code

import "testing"

/**
字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。



示例：

输入：S = "ababcbacadefegdehijhklij"
输出：[9,7,8]
解释：
划分结果为 "ababcbaca", "defegde", "hijhklij"。
每个字母最多出现在一个片段中。
像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。


提示：

S的长度在[1, 500]之间。
S只包含小写字母 'a' 到 'z' 。
*/
func Test763(t *testing.T) {
	t.Log(partitionLabels("ababcbacadefegdehijhklij")) // 9 7 8
}

// https://leetcode-cn.com/problems/partition-labels/solution/shou-hua-tu-jie-hua-fen-zi-mu-qu-jian-ji-lu-zui-yu/
func partitionLabels(S string) []int {
	maxPos := map[byte]int{}
	for i := 0; i < len(S); i++ {
		maxPos[S[i]] = i
	}

	res := []int{}
	start := 0
	scannedCharMaxPos := 0
	for i := 0; i < len(S); i++ {
		curCharMaxPos := maxPos[S[i]]
		if curCharMaxPos > scannedCharMaxPos {
			scannedCharMaxPos = curCharMaxPos
		}
		if i == scannedCharMaxPos {
			res = append(res, i-start+1)
			start = i + 1
		}
	}
	return res
}
