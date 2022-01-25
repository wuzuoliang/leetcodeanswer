package Code

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

/**
Given an unsorted array of integers, find the length of longest increasing subsequence.

Example:

[10,9,2,5,3,7,101,18]
[2,3,7,101]
4
Note:

There may be more than one LIS combination, it is only necessary for you to return the length.
Your algorithm should run in O(n2) complexity.
Follow up: Could you improve it to O(n log n) time complexity?
二分查找解法: https://mp.weixin.qq.com/s/7QFapCuvi-2nkh6gREcR9g
*/

func TestLongestIncreasingSubsequence(t *testing.T) {
	convey.Convey("TestLongestIncreasingSubsequence", t, func() {
		convey.Convey("10,9,2,5,3,7,101,18", func() {
			convey.So(IntShouldEqual(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}), 4), convey.ShouldBeTrue)
		})

		convey.Convey("1,9,5,9,3", func() {
			convey.So(IntShouldEqual(lengthOfLIS([]int{1, 9, 5, 9, 3}), 3), convey.ShouldBeTrue)
		})
	})
}

/**
我们分两种情况进行讨论：

如果nums[i]比前面的所有元素都小，那么dp[i]等于1（即它本身）（该结论正确）
如果nums[i]前面存在比他小的元素nums[j]，那么dp[i]就等于dp[j]+1（该结论错误，比如nums[3]>nums[0]，即9>1,但是dp[3]并不等于dp[0]+1）

我们先初步得出上面的结论，但是我们发现了一些问题。**因为dp[i]前面比他小的元素，不一定只有一个！**
可能除了 nums[j]，还包括 nums[k]，nums[p] 等等等等。所以 dp[i] 除了可能等于 dp[j]+1，还有可能等于 dp[k]+1，dp[p]+1 等等等等。所以我们求 dp[i]，需要找到 dp[j]+1，dp[k]+1，dp[p]+1 等等等等 中的最大值。（我在3个等等等等上都进行了加粗，主要是因为初学者非常容易在这里摔跟斗！这里强调的目的是希望能记住这道题型！） 即：
dp[i] = max(dp[j]+1，dp[k]+1，dp[p]+1，.....)
只要满足：

nums[i] > nums[j]
nums[i] > nums[k]
nums[i] > nums[p]
....

最后，我们只需要找到dp数组中的最大值，就是我们要找的答案。
复杂度分析： 时间复杂度O(n^2),空间复杂度O(n)
*/
func lengthOfLIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	result := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = Max(dp[j]+1, dp[i])
			}
		}
		result = Max(result, dp[i])
	}
	return result
}
