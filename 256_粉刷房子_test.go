package Code

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
假如有一排房子，共 n 个，每个房子可以被粉刷成红色、蓝色或者绿色这三种颜色中的一种，你需要粉刷所有的房子并且使其相邻的两个房子颜色不能相同。

当然，因为市场上不同颜色油漆的价格不同，所以房子粉刷成不同颜色的花费成本也是不同的。每个房子粉刷成不同颜色的花费是以一个n x 3的正整数矩阵 costs 来表示的。

例如，costs[0][0] 表示第 0 号房子粉刷成红色的成本花费；costs[1][2]表示第 1 号房子粉刷成绿色的花费，以此类推。

请计算出粉刷完所有房子最少的花费成本。



示例 1：

输入: costs = [[17,2,17],[16,16,5],[14,3,19]]
输出: 10
解释: 将 0 号房子粉刷成蓝色，1 号房子粉刷成绿色，2 号房子粉刷成蓝色。
    最少花费: 2 + 5 + 3 = 10。
示例 2：

输入: costs = [[7,6,2]]
输出: 2


提示:

costs.length == n
costs[i].length == 3
1 <= n <= 100
1 <= costs[i][j] <= 20

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/JEj789
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test256(t *testing.T) {
	convey.Convey("Test256", t, func() {
		convey.Convey("[[17,2,17],[16,16,5],[14,3,19]]", func() {
			convey.So(minCost([][]int{
				{17, 2, 17}, {16, 16, 5}, {14, 3, 19},
			}), convey.ShouldEqual, 10)
		})
		convey.Convey("[[7,6,2]]", func() {
			convey.So(minCost([][]int{
				{7, 6, 2},
			}), convey.ShouldEqual, 2)
		})
		convey.Convey("[[3,5,3],[6,17,6],[7,13,18],[9,10,18]]", func() {
			convey.So(minCost([][]int{
				{3, 5, 3},
				{6, 17, 6},
				{7, 13, 18},
				{9, 10, 18},
			}), convey.ShouldEqual, 26)
		})
	})
}

func minCost(costs [][]int) int {
	houses := len(costs)
	dps := make([][]int, houses)
	for i := 0; i < houses; i++ {
		dps[i] = make([]int, 3)
	}
	dps[0][0] = costs[0][0]
	dps[0][1] = costs[0][1]
	dps[0][2] = costs[0][2]

	for j := 1; j < houses; j++ {
		for i := 0; i < 3; i++ {
			switch i {
			case 0:
				dps[j][i] = Min(dps[j-1][1], dps[j-1][2]) + costs[j][i]
			case 1:
				dps[j][i] = Min(dps[j-1][0], dps[j-1][2]) + costs[j][i]
			case 2:
				dps[j][i] = Min(dps[j-1][1], dps[j-1][0]) + costs[j][i]
			}
		}
	}
	return Mins(dps[houses-1][0], dps[houses-1][1], dps[houses-1][2])
}
