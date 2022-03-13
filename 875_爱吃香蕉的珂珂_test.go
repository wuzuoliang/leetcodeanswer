package Code

import (
	"testing"
)

/**
珂珂喜欢吃香蕉。这里有N堆香蕉，第 i 堆中有piles[i]根香蕉。警卫已经离开了，将在H小时后回来。

珂珂可以决定她吃香蕉的速度K（单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。

珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。

返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。



示例 1：

输入: piles = [3,6,7,11], H = 8
输出: 4
示例2：

输入: piles = [30,11,23,4,20], H = 5
输出: 30
示例3：

输入: piles = [30,11,23,4,20], H = 6
输出: 23


提示：

1 <= piles.length <= 10^4
piles.length <= H <= 10^9
1 <= piles[i] <= 10^9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/koko-eating-bananas
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test875(t *testing.T) {
	t.Log(minEatingSpeed([]int{3, 6, 7, 11}, 8))       //4
	t.Log(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5)) //30
	t.Log(minEatingSpeed([]int{30, 11, 23, 4, 20}, 6)) // 23
}

func TestEatAll(t *testing.T) {
	t.Log(canEatAll([]int{3, 6, 7, 11}, 4))
	t.Log(canEatAll([]int{3, 6, 7, 11}, 3))
	t.Log(canEatAll([]int{30, 11, 23, 4, 20}, 29))
	t.Log(canEatAll([]int{30, 11, 23, 4, 20}, 30))
	t.Log(canEatAll([]int{30, 11, 23, 4, 20}, 23))
	t.Log(canEatAll([]int{30, 11, 23, 4, 20}, 22))
}

func minEatingSpeed(piles []int, h int) int {
	lowerSpeed := 1
	maxSpeed := 1000000001 // 这里可以取数组里的最大值
	for lowerSpeed < maxSpeed {
		midSpeed := lowerSpeed + (maxSpeed-lowerSpeed)/2
		if canEatAll(piles, midSpeed) == h {
			maxSpeed = midSpeed
		} else if canEatAll(piles, midSpeed) > h {
			lowerSpeed = midSpeed + 1
		} else if canEatAll(piles, midSpeed) < h {
			maxSpeed = midSpeed
		}
	}
	return lowerSpeed
}

func canEatAll(piles []int, speed int) int {
	hours := 0
	for i := 0; i < len(piles); i++ {
		hours += piles[i] / speed
		if piles[i]%speed > 0 {
			hours++
		}
	}
	return hours
}
