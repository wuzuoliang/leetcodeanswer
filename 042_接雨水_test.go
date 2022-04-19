package Code

import "testing"

/**
给定n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。



示例 1：



输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：

输入：height = [4,2,0,3,2,5]
输出：9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/trapping-rain-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test42(t *testing.T) {
	t.Log(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	t.Log(trap([]int{4, 2, 0, 3, 2, 5}))
}

// https://leetcode-cn.com/problems/trapping-rain-water/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-w-8/
func trap(height []int) int {
	// 时间复杂度O(n),空间复杂度O(n)
	// 表示该位置左右两边最高的第二大值
	//memo:=make([]int,len(height))
	//for i:=0;i<len(height);i++{
	//	leftMax:=height[i]
	//	rightMax:=height[i]
	//	for j:=0;j<len(height);j++{
	//		if height[j]>leftMax && j<=i{
	//			leftMax=height[j]
	//		}
	//		if height[j]>rightMax && j>i{
	//			rightMax=height[j]
	//		}
	//	}
	//	memo[i]=Min(leftMax,rightMax)
	//}
	//
	//maxArea:=0
	//for i:=0;i<len(height);i++{
	//	maxArea+=memo[i]-height[i]
	//}
	//return maxArea

	// 时间复杂度O(n),空间复杂度O(1)
	left := 0
	right := len(height) - 1
	leftMax := 0
	rightMax := 0
	maxArea := 0
	for left < right {
		// l_max是height[0..left]中最高柱子的高度，r_max是height[right..end]的最高柱子的高度
		leftMax = Max(leftMax, height[left])
		rightMax = Max(rightMax, height[right])

		if leftMax < rightMax {
			maxArea += leftMax - height[left]
			left++
		} else {
			maxArea += rightMax - height[right]
			right--
		}
	}
	return maxArea
}
