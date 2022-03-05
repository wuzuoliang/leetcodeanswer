package Code

import (
	"math"
	"testing"
)

/**
给定一个数组 prices ，它的第i 个元素prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。


示例 1：

输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
示例 2：

输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 没有交易完成, 所以最大利润为 0。

提示：

1 <= prices.length <= 105
0 <= prices[i] <= 104

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test121(t *testing.T) {
	t.Log(maxProfit121([]int{1, 2, 3, 5, 6}))    //5
	t.Log(maxProfit121([]int{7, 1, 5, 3, 6, 4})) //5
	t.Log(maxProfit121([]int{7, 6, 4, 3, 1}))    //0
}

/**
https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247494095&idx=4&sn=7aed55b22e93c0e43b83172923b51acc&scene=21#wechat_redirect
如果我们真的在买卖股票，我们肯定会想：如果我是在历史最低点买的股票就好了！
太好了，在题目中，我们只要用一个变量记录一个历史最低价格 minprice，我们就可以假设自己的股票是在那天买的。
那么我们在第 i 天卖出股票能得到的利润就是 prices[i] - minprice。
*/
func maxProfit121(prices []int) int {
	// 此方法相当于折线图寻找最低点，然后求比最低点后面的最高点和之前已有的最大差值做比较
	minprice := math.MaxInt
	maxprofit := 0
	for _, price := range prices {
		maxprofit = Max(maxprofit, price-minprice)
		minprice = Min(price, minprice)
	}
	return maxprofit
}

/**
k == 1
动态规划方法
// 原始版本
int maxProfit_k_1(int[] prices) {
    int n = prices.length;
    int[][] dp = new int[n][2];
    for (int i = 0; i < n; i++) {
        if (i - 1 == -1) {
            // base case
            dp[i][0] = 0;
            dp[i][1] = -prices[i];
            continue;
        }
        dp[i][0] = Math.max(dp[i-1][0], dp[i-1][1] + prices[i]);
        dp[i][1] = Math.max(dp[i-1][1], -prices[i]);
    }
    return dp[n - 1][0];
}

// 空间复杂度优化版本
int maxProfit_k_1(int[] prices) {
    int n = prices.length;
    // base case: dp[-1][0] = 0, dp[-1][1] = -infinity
    int dp_i_0 = 0, dp_i_1 = Integer.MIN_VALUE;
    for (int i = 0; i < n; i++) {
        // dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
        dp_i_0 = Math.max(dp_i_0, dp_i_1 + prices[i]);
        // dp[i][1] = max(dp[i-1][1], -prices[i])
        dp_i_1 = Math.max(dp_i_1, -prices[i]);
    }
    return dp_i_0;
}
*/
