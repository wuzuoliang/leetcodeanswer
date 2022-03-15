package Code

import "testing"

/**
给定两个字符串s1和s2，返回 使两个字符串相等所需删除字符的ASCII值的最小和。



示例 1:

输入: s1 = "sea", s2 = "eat"
输出: 231
解释: 在 "sea" 中删除 "s" 并将 "s" 的值(115)加入总和。
在 "eat" 中删除 "t" 并将 116 加入总和。
结束时，两个字符串相等，115 + 116 = 231 就是符合条件的最小和。
示例2:

输入: s1 = "delete", s2 = "leet"
输出: 403
解释: 在 "delete" 中删除 "dee" 字符串变成 "let"，
将 100[d]+101[e]+101[e] 加入总和。在 "leet" 中删除 "e" 将 101[e] 加入总和。
结束时，两个字符串都等于 "let"，结果即为 100+101+101+101 = 403 。
如果改为将两个字符串转换为 "lee" 或 "eet"，我们会得到 433 或 417 的结果，比答案更大。


提示:

0 <= s1.length, s2.length <= 1000
s1和s2由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-ascii-delete-sum-for-two-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test712(t *testing.T) {
	t.Log(minimumDeleteSum("sea", "eat"))
}

func minimumDeleteSum(s1 string, s2 string) int {
	t1 := len(s1)
	t2 := len(s2)
	mem := make([][]int, t1)
	for i := 0; i < t1; i++ {
		mem[i] = make([]int, t2)
		for j := 0; j < t2; j++ {
			mem[i][j] = -1
		}
	}
	return dp712(s1, 0, s2, 0, mem)
}

func dp712(text1 string, i1 int, text2 string, i2 int, mem [][]int) int {
	res := 0
	if len(text1) == i1 {
		for ; i2 < len(text2); i2++ {
			res += int(text2[i2])
		}
		return res
	}

	if len(text2) == i2 {
		for ; i1 < len(text1); i1++ {
			res += int(text1[i1])
		}
		return res
	}

	if mem[i1][i2] != -1 {
		return mem[i1][i2]
	}

	if text1[i1] == text2[i2] {
		mem[i1][i2] = dp712(text1, i1+1, text2, i2+1, mem)
		return mem[i1][i2]
	} else {
		mem[i1][i2] = Min(
			int(text1[i1])+dp712(text1, i1+1, text2, i2, mem),
			int(text2[i2])+dp712(text1, i1, text2, i2+1, mem))
		return mem[i1][i2]
	}
}

// 自底向上方式
func minimumDeleteSumUp(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	f := make([][]int, m+1)
	ans := 0
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		ans += int(s1[i])
	}
	for i := 0; i < n; i++ {
		ans += int(s2[i])
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				f[i][j] = f[i-1][j-1] + int(s1[i-1])
			} else {
				f[i][j] = Max(f[i-1][j], f[i][j-1])
			}
		}
	}
	return ans - 2*f[m][n]
}
