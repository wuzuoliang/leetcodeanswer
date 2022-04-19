package Code

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

/**
假设你正在爬楼梯。需要 n阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？



示例 1：

输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶
示例 2：

输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶


提示：

1 <= n <= 45

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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
