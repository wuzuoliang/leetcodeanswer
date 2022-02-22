package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
给定一个长度为 n 的整数数组height。有n条垂线，第 i 条线的两个端点是(i, 0)和(i, height[i])。

找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。

示例 1：

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为49。
示例 2：

输入：height = [1,1]
输出：1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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

// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247494095&idx=5&sn=8a69b2ca4d48e8b4db2b153a405c6e52&scene=21#wechat_redirect
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
	max := 0
	for i < j {
		if height[i] <= height[j] {
			max = Max(max, height[i]*(j-i))
			i++
		} else {
			max = Max(max, height[j]*(j-i))
			j--
		}
	}
	return max
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
