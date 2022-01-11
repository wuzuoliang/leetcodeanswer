package Code

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

/**
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

Note: You can only move either do1wn or right at any point in time.

Example:

Input:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
Output: 7
Explanation: Because the path 1→3→1→1→1 minimizes the sum.
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
	})
}

func minPathSum(grid [][]int) int {
	length := len(grid)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
	}
	if length == 1 {
		return grid[0][0]
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
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
	return dp[length-1][length-1]
}
