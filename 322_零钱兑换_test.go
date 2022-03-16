package Code

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回-1 。

你可以认为每种硬币的数量是无限的。



示例1：

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
示例 2：

输入：coins = [2], amount = 3
输出：-1
示例 3：

输入：coins = [1], amount = 0
输出：0


提示：

1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/coin-change
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test322(t *testing.T) {
	convey.Convey("coins = [1, 2, 5], amount = 11", t, func() {
		convey.So(coinChange([]int{1, 2, 5}, 11), convey.ShouldEqual, 3)
	})
	convey.Convey("coins = [2], amount = 3", t, func() {
		convey.So(coinChange([]int{2}, 3), convey.ShouldEqual, -1)
	})
}

// 为啥dp数组初始化为amount + 1呢，
// 因为凑成amount金额的硬币数最多只可能等于amount（全用 1 元面值的硬币），
// 所以初始化为amount + 1就相当于初始化为正无穷，便于后续取最小值。
func coinChange(ck []int, amount int) int {
	if amount < 0 {
		return -1
	}
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(ck); j++ {
			if i-ck[j] < 0 {
				continue
			}
			dp[i] = Min(dp[i], 1+dp[i-ck[j]])
		}
	}
	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}
