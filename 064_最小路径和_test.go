package Code

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

/**
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。

示例 1：
输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。

示例 2：
输入：grid = [[1,2,3],[4,5,6]]
输出：12

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func TestMinimumPathSum(t *testing.T) {
	convey.Convey("TestMinimumPathSum", t, func() {
		convey.Convey("[	[1,3,1],   [1,5,1],  [4,2,1]]", func() {
			convey.So(IntShouldEqual(minPathSum([][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			}), 7), convey.ShouldBeTrue)
		})
		convey.Convey("[	[1,2,3],   [4,5,6]]", func() {
			convey.So(IntShouldEqual(minPathSum([][]int{
				{1, 2, 3},
				{4, 5, 6},
			}), 12), convey.ShouldBeTrue)
		})
		convey.Convey("[	[9,1,4,8]]", func() {
			convey.So(IntShouldEqual(minPathSum([][]int{
				{9, 1, 4, 8},
			}), 22), convey.ShouldBeTrue)
		})
	})
}

func minPathSum(grid [][]int) int {
	length := len(grid)
	dp := make([][]int, length)

	for i := 0; i < length; i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < length; i++ {
		for j := 0; j < len(dp[i]); j++ {
			if j == 0 {
				if i == 0 {
					dp[i][j] = dp[0][0]
				} else {
					dp[i][j] = dp[i-1][j] + grid[i][j]
				}
			} else if i == 0 {
				if j == 0 {
					dp[i][j] = dp[0][0]
				} else {
					dp[i][j] = dp[0][j-1] + grid[i][j]
				}
			} else {
				dp[i][j] = Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[length-1][len(dp[length-1])-1]
}
