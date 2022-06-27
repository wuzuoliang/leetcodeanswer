package Code

import "testing"

/**
给定字符串列表strs ，返回其中 最长的特殊序列 。如果最长特殊序列不存在，返回 -1 。

特殊序列 定义如下：该序列为某字符串 独有的子序列（即不能是其他字符串的子序列）。

s的子序列可以通过删去字符串s中的某些字符实现。

例如，"abc"是 "aebdc"的子序列，因为您可以删除"aebdc"中的下划线字符来得到 "abc"。"aebdc"的子序列还包括"aebdc"、 "aeb"和 ""(空字符串)。


示例 1：

输入: strs = ["aba","cdc","eae"]
输出: 3
示例 2:

输入: strs = ["aaa","aaa","aa"]
输出: -1


提示:

2 <= strs.length <= 50
1 <= strs[i].length <= 10
strs[i]只包含小写英文字母


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-uncommon-subsequence-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test522(t *testing.T) {
	t.Log(findLUSlength([]string{"aba", "cdc", "eae"})) // 3
	t.Log(findLUSlength([]string{"aaa", "aa", "a"}))    // -1
}

func isSubseq(s, t string) bool {
	ptS := 0
	for ptT := range t {
		if s[ptS] == t[ptT] {
			if ptS++; ptS == len(s) {
				return true
			}
		}
	}
	return false
}

func findLUSlength(strs []string) int {
	ans := -1
next:
	for i, s := range strs {
		for j, t := range strs {
			if i != j && isSubseq(s, t) {
				continue next
			}
		}
		if len(s) > ans {
			ans = len(s)
		}
	}
	return ans
}
