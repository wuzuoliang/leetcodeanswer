package Code

import "testing"

/**
给定两个字符串text1 和text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。

一个字符串的子序列是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。



示例 1：

输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace" ，它的长度为 3 。
示例 2：

输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc" ，它的长度为 3 。
示例 3：

输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0 。


提示：

1 <= text1.length, text2.length <= 1000
text1 和text2 仅由小写英文字符组成。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test1143(t *testing.T) {
	t.Log(longestCommonSubsequenceUP("abcde", "ace")) //3
	t.Log(longestCommonSubsequenceDown("abc", "abc")) //3
	t.Log(longestCommonSubsequenceDown("abc", "def")) //0
}

// 自底向上
func longestCommonSubsequenceUP(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = Max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}

// longestCommonSubsequenceDown 自顶向下
func longestCommonSubsequenceDown(text1 string, text2 string) int {
	t1 := len(text1)
	t2 := len(text2)
	mem := make([][]int, t1)
	for i := 0; i < t1; i++ {
		mem[i] = make([]int, t2)
		for j := 0; j < t2; j++ {
			mem[i][j] = -1
		}
	}
	return dp1143(text1, 0, text2, 0, mem)
}

func dp1143(text1 string, i1 int, text2 string, i2 int, mem [][]int) int {
	if len(text1) == i1 || len(text2) == i2 {
		return 0
	}

	if mem[i1][i2] != -1 {
		return mem[i1][i2]
	}

	if text1[i1] == text2[i2] {
		mem[i1][i2] = 1 + dp1143(text1, i1+1, text2, i2+1, mem)
		return mem[i1][i2]
	} else {
		mem[i1][i2] = Max(
			dp1143(text1, i1+1, text2, i2, mem),
			dp1143(text1, i1, text2, i2+1, mem))
		return mem[i1][i2]
	}
}
