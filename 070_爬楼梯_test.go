package Code

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

/**
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
*/

func Test_climbStairs(t *testing.T) {
	convey.Convey("Test_climbStairs", t, func() {
		convey.Convey("1", func() {
			convey.So(IntShouldEqual(climbStairs(1), 1), convey.ShouldBeTrue)
		})
		convey.Convey("2", func() {
			convey.So(IntShouldEqual(climbStairs(2), 2), convey.ShouldBeTrue)
		})
		convey.Convey("3", func() {
			convey.So(IntShouldEqual(climbStairs(3), 3), convey.ShouldBeTrue)
		})
		convey.Convey("4", func() {
			convey.So(IntShouldEqual(climbStairs(4), 5), convey.ShouldBeTrue)
		})
		convey.Convey("5", func() {
			convey.So(IntShouldEqual(climbStairs(5), 8), convey.ShouldBeTrue)
		})
		convey.Convey("6", func() {
			convey.So(IntShouldEqual(climbStairs(6), 13), convey.ShouldBeTrue)
		})
	})
}

/**
跟斐波那契数列非常相似，假设梯子有n层，那么如何爬到第n层呢，因为每次只能爬1或2步，那么爬到第n层的方法要么是从第 n-1 层一步上来的，要不就是从 n-2 层2步上来的，所以递推公式非常容易的就得出了：dp[n] = dp[n-1] + dp[n-2]
*/
func climbStairs(n int) int {
	a := 1
	b := 1
	for n >= 1 {
		n--
		b += a
		a = b - a
	}
	return a

	// res := make([]int, n+1)
	// res[0] = 1
	// res[1] = 1
	// for i := 2; i <= n; i++ {
	// 	res[i] = res[i-1] + res[i-2]
	// }
	// return res[n]
}

/**
一步一个台阶，两个台阶，三个台阶，.......，直到 m个台阶。问有多少种不同的方法可以爬到楼顶呢？

func climbStairs(n int) int {
	//定义
	dp := make([]int, n+1)
	//初始化
	dp[0] = 1
	// 本题物品只有两个1,2
	m := 2
	// 遍历顺序
	for j := 1; j <= n; j++ {	//先遍历背包
		for i := 1; i <= m; i++ {	//再遍历物品
			if j >= i {
				dp[j] += dp[j-i]
			}
			//fmt.Println(dp)
		}
	}
	return dp[n]
}
*/
