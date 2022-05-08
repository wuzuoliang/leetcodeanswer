package Code

import "testing"

/**
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

回文串 是正着读和反着读都一样的字符串。



示例 1：

输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]
示例 2：

输入：s = "a"
输出：[["a"]]


提示：

1 <= s.length <= 16
s 仅由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-partitioning
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test131(t *testing.T) {
	t.Log(partition131("aab"))
	t.Log(partition131("a"))
}

func partition131(s string) [][]string {
	res := make([][]string, 0)
	resOne := make([]string, 0)
	var dfs func(res *[][]string, s string, start int, curstr []string)
	dfs = func(res *[][]string, s string, start int, curstr []string) {
		if start == len(s) {
			tmp := make([]string, len(curstr))
			copy(tmp, curstr)
			*res = append(*res, tmp)
			return
		}

		for index := start; index < len(s); index++ {
			if !isPartition(s, start, index) {
				continue
			}
			pStr := s[start : index+1]
			curstr = append(curstr, pStr)
			dfs(res, s, index+1, curstr)
			curstr = curstr[:len(curstr)-1]
		}
	}
	dfs(&res, s, 0, resOne)
	return res
}

//判断是否为回文
func isPartition(s string, startIndex, end int) bool {
	left := startIndex
	right := end
	for left < right {
		if s[left] != s[right] {
			return false
		}
		//移动左右指针
		left++
		right--
	}
	return true
}
