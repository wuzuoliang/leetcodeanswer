package Code

import (
	"math"
	"testing"
)

/**
给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。

请你找出符合题意的 最短 子数组，并输出它的长度。



示例 1：

输入：nums = [2,6,4,8,10,9,15]
输出：5
解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
示例 2：

输入：nums = [1,2,3,4]
输出：0
示例 3：

输入：nums = [1]
输出：0


提示：

1 <= nums.length <= 104
-105 <= nums[i] <= 105


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test581(t *testing.T) {
	t.Log(findUnsortedSubarray([]int{1, 2, 3, 4}))            //0
	t.Log(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15})) //5
}

// 只需要知道替换地方的最小ID，和最大ID就可以知道范围了
func findUnsortedSubarrayBest(nums []int) int {
	min := nums[len(nums)-1]
	max := nums[0]
	start, end := 0, -1

	for i := 0; i < len(nums); i++ {
		if nums[i] >= max { //从左到右维持最大值，寻找右边界end
			max = nums[i]
		} else {
			end = i
		}

		if nums[len(nums)-i-1] <= min { //从右到左维持最小值，寻找左边界begin
			min = nums[len(nums)-i-1]
		} else {
			start = len(nums) - i - 1
		}
	}
	return end - start + 1
}

func findUnsortedSubarray(nums []int) int {
	minInx := math.MaxInt
	maxInx := math.MinInt
	sortCount := 0
	uperStack := make([]int, len(nums))
	stackLen := 0
	for i := 0; i < len(nums); i++ {
		if stackLen == 0 {
			uperStack[i] = nums[i]
			stackLen++
			continue
		}
		if nums[i] >= uperStack[stackLen-1] {
			uperStack[i] = nums[i]
			stackLen++
			continue
		}

		uperStack[i] = nums[i]
		stackLen++
		for j := stackLen - 1; j > 0; j-- {
			if uperStack[j] < uperStack[j-1] {
				uperStack[j-1], uperStack[j] = uperStack[j], uperStack[j-1]
				minInx = Min(j-1, minInx)
				maxInx = Max(minInx, stackLen-1)
			}
		}
		sortCount++
	}
	if sortCount > 0 {
		return maxInx - minInx + 1
	} else {
		return 0
	}
}
