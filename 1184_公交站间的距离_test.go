package Code

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
环形公交路线上有n个站，按次序从0到n - 1进行编号。我们已知每一对相邻公交站之间的距离，distance[i]表示编号为i的车站和编号为(i + 1) % n的车站之间的距离。

环线上的公交车都可以按顺时针和逆时针的方向行驶。

返回乘客从出发点start到目的地destination之间的最短距离。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/distance-between-bus-stops
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test1184(t *testing.T) {
	convey.Convey("Test1184", t, func() {
		convey.Convey("distance = [1,2,3,4], start = 0, destination = 1,expected 1", func() {
			convey.So(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 1), convey.ShouldEqual, 1)
		})
		convey.Convey("distance = [1,2,3,4], start = 0, destination = 2,expected 3", func() {
			convey.So(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 2), convey.ShouldEqual, 3)
		})
		convey.Convey("distance = [1,2,3,4], start = 0, destination = 3,expected 4", func() {
			convey.So(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 3), convey.ShouldEqual, 4)
		})
	})
}

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	dis := 0
	total := 0
	for _, v := range distance {
		total += v
	}

	if destination == start {
		return 0
	} else if destination > start {
		for i := start; i < destination; i++ {
			dis += distance[i]
		}
	} else {
		for i := destination; i < start; i++ {
			dis += distance[i]
		}
	}
	return Min(dis, total-dis)
}
