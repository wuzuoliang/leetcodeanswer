package Code

import "testing"

/**
You are given coins of different denominations and a total amount of money. Write a function to compute the number of combinations that make up that amount. You may assume that you have infinite number of each kind of coin.

Note: You can assume that

0 <= amount <= 5000
1 <= coin <= 5000
the number of coins is less than 500
the answer is guaranteed to fit into signed 32-bit integer


Example 1:

Input: amount = 5, coins = [1, 2, 5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1


Example 2:

Input: amount = 3, coins = [2]
Output: 0
Explanation: the amount of 3 cannot be made up just with coins of 2.


Example 3:

Input: amount = 10, coins = [10]
Output: 1

*/
func TestCoinChange2(t *testing.T) {
	t.Log(coinChange2([]int{1, 2, 5}, 5))
	t.Log(coinChange2([]int{2}, 3))
	t.Log(coinChange2([]int{10}, 10))
}

func coinChange2(coins []int, amount int) int {
	if amount == 0 {
		return 1
	}
	if len(coins) == 0 {
		return 0
	}

	n := len(coins)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				// 可以装下，选择不装  dp[i-1][j] ，或者选择装 dp[i][j-coins[i-1]]，两类之和
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {
				// 不可以装下
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][amount]
}

/**
我们通过观察可以发现，dp数组的转移只和dp[i][..]和dp[i-1][..]有关，所以可以压缩状态，进一步降低算法的空间复杂度：

int change(int amount, int[] coins) {
    int n = coins.length;
    int[] dp = new int[amount + 1];
    dp[0] = 1; // base case
    for (int i = 0; i < n; i++)
        for (int j = 1; j <= amount; j++)
            if (j - coins[i] >= 0)
                dp[j] = dp[j] + dp[j-coins[i]];

    return dp[amount];
}
*/
