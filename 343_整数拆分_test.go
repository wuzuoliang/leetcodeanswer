package Code

import "testing"

/**
给定一个正整数n，将其拆分为 k 个 正整数 的和（k >= 2），并使这些整数的乘积最大化。

返回 你可以获得的最大乘积。



示例 1:

输入: n = 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。
示例2:

输入: n = 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 ×3 ×4 = 36。


提示:

2 <= n <= 58

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-break
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test343(t *testing.T) {
	t.Log(integerBreak(10))
}

func integerBreak(n int) int {
	/**
	  动态五部曲
	  1.确定dp下标及其含义
	  2.确定递推公式
	  3.确定dp初始化
	  4.确定遍历顺序
	  5.打印dp
	  **/
	dp := make([]int, n+1)
	dp[1] = 0
	dp[2] = 1
	for i := 3; i < n+1; i++ {
		for j := 1; j < i-1; j++ {
			// i可以差分为i-j和j。由于需要最大值，
			// 故需要通过j遍历所有存在的值，
			// 取其中最大的值作为当前i的最大值，
			// 在求最大值的时候，
			// 一个是j与i-j相乘，
			// 一个是j与dp[i-j].
			dp[i] = Max(dp[i], Max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}
