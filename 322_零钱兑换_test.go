package Code

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
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
func coinChange(coins []int, amount int) int {
	// 先遍历物品,再遍历背包
	dp := make([]int, amount+1)
	// 初始化dp[0]
	dp[0] = 0
	// 初始化为math.MaxInt32
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt32
	}

	// 遍历物品
	for i := 0; i < len(coins); i++ {
		// 遍历背包
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				// 推导公式
				dp[j] = Min(dp[j], dp[j-coins[i]]+1)
				//fmt.Println(dp,j,i)
			}
		}
	}
	// 没找到能装满背包的, 就返回-1
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
