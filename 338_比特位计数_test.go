package Code

import "testing"

/**
给你一个整数 n ，对于0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。



示例 1：

输入：n = 2
输出：[0,1,1]
解释：
0 --> 0
1 --> 1
2 --> 10
示例 2：

输入：n = 5
输出：[0,1,1,2,1,2]
解释：
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/counting-bits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test338(t *testing.T) {
	t.Log(countBits(5))
	t.Log(countBitsDP(8))
}

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

func countBitsDP(n int) []int {
	if n == 0 {
		return []int{0}
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	c := 2
	for i := 2; i <= n; i++ {
		if i >= c*2 {
			c = c * 2
		}
		if i == c {
			dp[i] = 1
		} else {
			dp[i] = dp[i-c] + 1
		}
	}
	return dp
}
