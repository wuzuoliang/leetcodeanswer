package Code

import "testing"

/**
给你一个整数数组 nums 和一个整数 target 。

向数组中的每个整数前添加'+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：

例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。



示例 1：

输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
示例 2：

输入：nums = [1], target = 1
输出：1


提示：

1 <= nums.length <= 20
0 <= nums[i] <= 1000
0 <= sum(nums[i]) <= 1000
-1000 <= target <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/target-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test494(t *testing.T) {
	t.Log(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3)) //5
	t.Log(findTargetSumWays([]int{1}, 1))             //1
}

func findTargetSumWays(nums []int, target int) (count int) {
	var backtrack func(int, int)
	backtrack = func(index, sum int) {
		if index == len(nums) {
			if sum == target {
				count++
			}
			return
		}
		backtrack(index+1, sum+nums[index])
		backtrack(index+1, sum-nums[index])
	}
	backtrack(0, 0)
	return
}

func findTargetSumWaysDP(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	diff := sum - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	n, neg := len(nums), diff/2
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1
	for i, num := range nums {
		for j := 0; j <= neg; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= num {
				dp[i+1][j] += dp[i][j-num]
			}
		}
	}
	return dp[n][neg]
}
