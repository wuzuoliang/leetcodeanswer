package Code

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

/**
第122题：买卖股票的最佳时机 II
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
在每一天，你可能会决定购买和/或出售股票。你在任何时候最多只能持有 一股 股票。你也可以购买它，然后在 同一天 出售。
返回 你能获得的 最大 利润。

示例 1:
输入: prices = [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
    随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。

示例 2:
输入: prices = [1,2,3,4,5]
输出: 4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
    注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
示例3:

输入: prices = [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestBestTimeToBuyAndSellStock2(t *testing.T) {
	convey.Convey("TestBestTimeToBuyAndSellStock2", t, func() {
		convey.Convey("7,1,5,3,6,4", func() {
			convey.So(IntShouldEqual(7, maxProfit122([]int{7, 1, 5, 3, 6, 4})), convey.ShouldBeTrue)
		})
		convey.Convey("1,2,3,4,5", func() {
			convey.So(IntShouldEqual(4, maxProfit122([]int{1, 2, 3, 4, 5})), convey.ShouldBeTrue)
		})
		convey.Convey("7,6,4,3,1", func() {
			convey.So(IntShouldEqual(0, maxProfit122([]int{7, 6, 4, 3, 1})), convey.ShouldBeTrue)
		})
		convey.Convey("1,6,5,7", func() {
			convey.So(IntShouldEqual(7, maxProfit122([]int{1, 6, 5, 7})), convey.ShouldBeTrue)
		})
	})
}

// 相当于查分数组的思想
func maxProfit122(prices []int) int {
	res := 0
	for i := 0; i < len(prices)-1; i++ {
		incr := prices[i+1] - prices[i]
		if incr > 0 {
			res += incr
		}
	}
	return res

	// 贪心算法
	//ans:=0
	//for i := 1; i < len(prices); i++ {
	//	ans += Max(0, prices[i]-prices[i-1])
	//}
	//return ans

}

/**
k = +infinity
// 原始版本
int maxProfit_k_inf(int[] prices) {
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
        dp[i][1] = Math.max(dp[i-1][1], dp[i-1][0] - prices[i]);
    }
    return dp[n - 1][0];
}

// 空间复杂度优化版本
int maxProfit_k_inf(int[] prices) {
    int n = prices.length;
    int dp_i_0 = 0, dp_i_1 = Integer.MIN_VALUE;
    for (int i = 0; i < n; i++) {
        int temp = dp_i_0;
        dp_i_0 = Math.max(dp_i_0, dp_i_1 + prices[i]);
        dp_i_1 = Math.max(dp_i_1, temp - prices[i]);
    }
    return dp_i_0;
}
*/
