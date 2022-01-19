package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

/**
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.
Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49
*/

func Test_maxArea(t *testing.T) {
	Convey("Test_maxArea", t, func() {

		Convey("[7,2,8]", func() {
			So(IntShouldEqual(maxArea([]int{7, 2, 8}), 14), ShouldBeTrue)
		})

		Convey("[1,8,6,2,5,4,8,3,7]", func() {
			So(IntShouldEqual(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49), ShouldBeTrue)
		})
	})
}
func maxArea(height []int) int {
	// O(n^n)
	//lens := len(height)
	//if lens < 2 {
	//	return 0
	//}
	//max := 0
	//for i := 0; i < lens; i++ {
	//	if i == 0 {
	//		calculateAreaRight(height, i, lens, &max)
	//	} else if i == lens-1 {
	//		calculateAreaLeft(height, i, lens, &max)
	//	} else {
	//		calculateAreaRight(height, i, lens, &max)
	//		calculateAreaLeft(height, i, lens, &max)
	//	}
	//}
	//return max

	// O(n)
	i, j := 0, len(height)-1
	max := float64(0)
	for i < j {
		if height[i] <= height[j] {
			max = math.Max(max, float64(height[i]*(j-i)))
			i++
		} else {
			max = math.Max(max, float64(height[j]*(j-i)))
			j--
		}
	}
	return int(max)
}

func calculateAreaRight(height []int, cur, lens int, area *int) {
	for j := cur + 1; j < lens; j++ {
		if height[j] >= height[cur] {
			*area = Max(height[cur]*(j-cur), *area)
		}
	}
}
func calculateAreaLeft(height []int, cur, lens int, area *int) {
	for j := cur - 1; j >= 0; j-- {
		if height[j] >= height[cur] {
			*area = Max(height[cur]*(cur-j), *area)
		}
	}
}
