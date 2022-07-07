package Code

import (
	"testing"
)

/**
给定一个包含大写字母和小写字母的字符串s，返回通过这些字母构造成的 最长的回文串。

在构造过程中，请注意 区分大小写 。比如"Aa"不能当做一个回文字符串。



示例 1:

输入:s = "abccccdd"
输出:7
解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
示例 2:

输入:s = "a"
输入:1
示例 3:

输入:s = "bb"
输入: 2


提示:

1 <= s.length <= 2000
s只能由小写和/或大写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test409(t *testing.T) {
	t.Log(longestPalindromeNum("abccccdd"))
	t.Log(longestPalindromeNum("a"))
	t.Log(longestPalindromeNum("bb"))
}

func longestPalindromeNum(s string) int {
	existMap := make(map[int]bool)
	count := 0
	for _, v := range s {
		idx := byte(v) - 'A'
		if _, ok := existMap[int(idx)]; ok {
			count += 2
			delete(existMap, int(idx))
		} else {
			existMap[int(idx)] = true
		}
	}
	// 上面是出现偶数次字数的个数
	if count < len(s) { // 至少可以放一个中间位置
		return count + 1
	} else {
		return count
	}
}
