package Code

import "testing"

/**
给定一个整数数组prices ，它的第 i 个元素prices[i] 是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1：

输入：k = 2, prices = [2,4,1]
输出：2
解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
示例 2：

输入：k = 2, prices = [3,2,6,5,0,3]
输出：7
解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。

提示：
0 <= k <= 100
0 <= prices.length <= 1000
0 <= prices[i] <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test188(t *testing.T) {
	t.Log(maxProfit188(2, []int{2, 4, 1}))          //2
	t.Log(maxProfit188(2, []int{3, 2, 6, 5, 0, 3})) //7
}

func maxProfit188(maxK int, prices []int) int {
	n := len(prices)
	if maxK > n/2 {
		return maxProfit121(prices)
	}

	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, maxK+1)
		for j := 0; j < maxK+1; j++ {
			dp[i][j] = []int{0, 1}
		}
	}

	for i := 0; i < n; i++ {
		for k := maxK; k >= 1; k-- {
			if i-1 == -1 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = Max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = Max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[n-1][maxK][0]
}

/**
k = any integer
若传入的k值会非常大，dp数组太大了，oom
*/
func maxProfitBest(k int, prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	if k <= 0 {
		return 0
	}
	// 贪心: 可视为不限买卖次数
	if k >= len(prices)/2 {
		var ans int
		for i := 0; i < len(prices)-1; i++ {
			if prices[i] < prices[i+1] {
				ans += prices[i+1] - prices[i]
			}
		}
		return ans
	}
	// 贪心：仅一次，则每次判断跟前面的最小值的收益即可
	if k == 1 {
		var ans int
		min := prices[0]
		for i := 0; i < len(prices); i++ {
			if ans < prices[i]-min {
				ans = prices[i] - min
			}
			if prices[i] < min {
				min = prices[i]
			}
		}
		return ans
	}
	// 基本状态：买入、卖出，细分第1~k次买入、第1~k次卖出
	// 由于仅依赖前一次状态，使用滚动数组，dp[0]表示前一次的状态
	dpBuy := make([][]int, 2, 2)
	dpSell := make([][]int, 2, 2)
	for i := 0; i < 2; i++ {
		dpBuy[i], dpSell[i] = make([]int, k), make([]int, k)
	}
	for i := 0; i < k; i++ {
		dpBuy[0][i] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		// 第一次买入需要特殊处理，取[0~i]的最小值
		dpBuy[1][0] = Max(dpBuy[0][0], -prices[i])
		dpSell[1][0] = Max(dpSell[0][0], dpBuy[0][0]+prices[i])
		//
		dpBuy[0][0], dpSell[0][0] = dpBuy[1][0], dpSell[1][0]
		for j := 1; j < k; j++ {
			// 当前第j次买入：1、前一次就是第j次买入  2、前一次为第j-1次卖出
			dpBuy[1][j] = Max(dpBuy[0][j], dpSell[0][j-1]-prices[i])
			// 当前第j次卖出：1、前一次就是第j次卖出  2、前一次为第j次买入
			dpSell[1][j] = Max(dpSell[0][j], dpBuy[0][j]+prices[i])
			//
			dpBuy[0][j], dpSell[0][j] = dpBuy[1][j], dpSell[1][j]
		}
	}
	return dpSell[1][k-1]
}
