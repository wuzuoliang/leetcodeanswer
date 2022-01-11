package Code

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestOneZeroPackage(t *testing.T) {
	convey.Convey("TestOneZeroPackage", t, func() {

		convey.Convey("2 3 4 5 - 3,4,5,6", func() {
			convey.So(
				oneZeroPackage([]zeroPackage{
					{2, 3}, {3, 4}, {4, 5}, {5, 6},
				}, 8), convey.ShouldEqual, 10)
		})
	})
}

type zeroPackage struct {
	Vol int
	Val int
}

func oneZeroPackage(goods []zeroPackage, size int) int {
	tmpGoods := make([]zeroPackage, 0)
	tmpGoods = append(tmpGoods, zeroPackage{0, 0})
	tmpGoods = append(tmpGoods, goods...)
	total := len(tmpGoods)
	dp := make([][]int, total)
	for i := 0; i < total; i++ {
		dp[i] = make([]int, size+1)
	}
	for i := 0; i < total; i++ {
		for j := 0; j <= size; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
				continue
			}
			withOut := dp[i-1][j]
			with := 0
			if j >= tmpGoods[i].Vol {
				with = tmpGoods[i].Val + dp[i-1][j-tmpGoods[i].Vol]
			}
			dp[i][j] = Max(with, withOut)
		}
	}
	return dp[total-1][size]
}
