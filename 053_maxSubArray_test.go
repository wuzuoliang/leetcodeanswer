package Code

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

/**
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/

func Test_maxSubArray(t *testing.T) {
	convey.Convey("Test_maxSubArray", t, func() {
		convey.So(IntShouldEqual(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), 6), convey.ShouldBeTrue)

		convey.So(IntShouldEqual(maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), 6), convey.ShouldBeTrue)
	})
}

// 复杂度分析：时间复杂度：O(N)。空间复杂度：O(N)
func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	result := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		result = max(dp[i], result)
	}
	return result
}

// O(nlogn)
func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return spilts2(nums, 0, len(nums)-1)
}

func spilts2(nums []int, left, right int) int {
	if left >= right {
		return nums[left]
	}
	mid := left + (right-left)/2
	lmax := spilts2(nums, left, mid-1)
	rmax := spilts2(nums, mid+1, right)
	mmax := nums[mid]
	t := mmax
	for i := mid - 1; i >= left; i-- {
		t += nums[i]
		mmax = Max(mmax, t)
	}
	t = mmax
	for i := mid + 1; i <= right; i++ {
		t += nums[i]
		mmax = Max(mmax, t)
	}
	return Max(mmax, Max(lmax, rmax))
}
