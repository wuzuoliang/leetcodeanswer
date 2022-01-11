package Code

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

/**
第122题：买卖股票的最佳时机 II
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。注意你不能在买入股票前卖出股票。

示例 1:

输入: [7,1,5,3,6,4]
输出: 7
1
2
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。

随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。

示例 2:

输入: [1,2,3,4,5]
输出: 4
1
2
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。

​注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。

因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。

示例 3:

输入: [7,6,4,3,1]
输出: 0
1
2
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
*/

func TestBestTimeToBuyAndSellStock2(t *testing.T) {
	convey.Convey("TestBestTimeToBuyAndSellStock2", t, func() {
		convey.Convey("7,1,5,3,6,4", func() {
			convey.So(IntShouldEqual(7, maxProfit([]int{7, 1, 5, 3, 6, 4})), convey.ShouldBeTrue)
		})
		convey.Convey("1,2,3,4,5", func() {
			convey.So(IntShouldEqual(4, maxProfit([]int{1, 2, 3, 4, 5})), convey.ShouldBeTrue)
		})
		convey.Convey("7,6,4,3,1", func() {
			convey.So(IntShouldEqual(0, maxProfit2([]int{7, 6, 4, 3, 1})), convey.ShouldBeTrue)
		})
	})
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[len(prices)-1][0]
}

func maxProfit2(prices []int) int {
	res := 0
	for i := 0; i < len(prices)-1; i++ {
		incr := prices[i+1] - prices[i]
		if incr > 0 {
			res += incr
		}
	}
	return res
}
