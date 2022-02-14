package Code

import "testing"

/**
给你输入一个数组nums和一个正整数k，请你判断nums是否能够被平分为元素和相同的k个子集。
*/

func Test698(t *testing.T) {
	t.Log(canPartitionKSubsets([]int{1, 2, 3, 4}, 2))
	t.Log(canPartitionKSubsets([]int{1, 2, 3, 4, 5, 6}, 3))
}

func canPartitionKSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	used := make([]bool, len(nums))
	target := sum / k
	return backtrack(k, nums, 0, used, 0, target)
}
func backtrack(k int, nums []int, start int, used []bool, curSum, target int) bool {
	// base case
	if k == 0 {
		// 所有桶都被装满了，而且 nums 一定全部用完了
		// 因为 target == sum / k
		return true
	}
	if curSum == target {
		return backtrack(k-1, nums, 0, used, 0, target)
	}

	for i := start; i < len(nums); i++ {

		if used[i] {
			continue
		}

		if curSum+nums[i] > target {
			continue
		}

		curSum += nums[i]
		used[i] = true
		if backtrack(k, nums, i+1, used, curSum, target) {
			return true
		}
		used[i] = false
		curSum -= nums[i]
	}
	return false
}
