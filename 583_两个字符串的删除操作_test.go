package Code

import "testing"

/**
给定两个单词word1和word2，返回使得word1和word2相同所需的最小步数。

每步可以删除任意一个字符串中的一个字符。



示例 1：

输入: word1 = "sea", word2 = "eat"
输出: 2
解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"
示例 2:

输入：word1 = "leetcode", word2 = "etco"
输出：4


提示：

1 <= word1.length, word2.length <= 500
word1和word2只包含小写英文字母

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/delete-operation-for-two-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test583(t *testing.T) {
	t.Log(minDistance("sea", "eat"))
}

// 题目让我们计算将两个字符串变得相同的最少删除次数，那我们可以思考一下，最后这两个字符串会被删成什么样子？
// 删除的结果不就是它俩的最长公共子序列嘛！
// 那么，要计算删除的次数，就可以通过最长公共子序列的长度推导出来
// 参考1143，求最长公共子序列
func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	lcs := longestCommonSubsequenceUP(word1, word2)
	return m + n - 2*lcs
}
