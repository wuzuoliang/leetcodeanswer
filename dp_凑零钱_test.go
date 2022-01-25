package Code

import (
	"testing"
)

/**
给你k种面值的硬币，面值分别为c1, c2 ... ck，每种硬币的数量无限，再给一个总金额amount，问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1 。
*/

func TestMakeTheChange(t *testing.T) {
	t.Log(coinChange([]int{1, 2, 5}, 11))
}

// 为啥dp数组初始化为amount + 1呢，因为凑成amount金额的硬币数最多只可能等于amount（全用 1 元面值的硬币），所以初始化为amount + 1就相当于初始化为正无穷，便于后续取最小值。
func coinChange(ck []int, amount int) int {
	if amount < 0 {
		return -1
	}
	dp := make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		dp[i] = i
		for j := 0; j < len(ck); j++ {
			if i-ck[j] < 0 {
				continue
			}
			dp[i] = Min(dp[i], 1+dp[i-ck[j]])
		}
	}
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}
