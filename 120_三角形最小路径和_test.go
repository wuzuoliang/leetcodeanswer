package Code

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

/**
Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

For example, given the following triangle

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]


The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

Note:
Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.
*/
func TestTriangle(t *testing.T) {
	convey.Convey("TestTriangle", t, func() {
		convey.Convey("[	[2],   [3,4],  [6,5,7], [4,1,8,3]]", func() {
			convey.So(IntShouldEqual(minimumTotal([][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			}), 11), convey.ShouldBeTrue)
		})
	})
}

func minimumTotal(triangle [][]int) int {
	high := len(triangle)
	if high == 1 {
		return triangle[0][0]
	}
	dp := make([][]int, high)
	for i := 0; i < high; i++ {
		dp[i] = make([]int, i+1)
	}
	dp[0][0] = triangle[0][0]
	dp[1][0] = dp[0][0] + triangle[1][0]
	dp[1][1] = dp[0][0] + triangle[1][1]

	for level := 1; level < high; level++ {
		for col := 0; col <= level; col++ {
			if col == 0 {
				dp[level][col] = dp[level-1][0] + triangle[level][col]
			} else if col == level {
				dp[level][col] = dp[level-1][level-1] + triangle[level][col]
			} else {
				dp[level][col] = (Min(dp[level-1][level-2], dp[level-1][level-1])) + triangle[level][col]
			}
		}
	}
	min := dp[high-1][0]
	for i := 1; i < high; i++ {
		min = Min(dp[high-1][i], min)
	}
	return min
}
