package Code

import "testing"

/**
给你一个非负整数数组 nums ，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

假设你总是可以到达数组的最后一个位置。



示例 1:

输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
示例 2:

输入: nums = [2,3,0,1,4]
输出: 2


提示:

1 <= nums.length <= 104
0 <= nums[i] <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jump-game-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test45(t *testing.T) {
	t.Log(jump([]int{2, 3, 0, 1, 4}))
	t.Log(jumpdp([]int{2, 3, 1, 1, 4}))
}

/**
我们不需要「递归地」计算出所有选择的具体结果然后比较求最值，而只需要做出那个最有「潜力」，看起来最优的选择即可。
i和end标记了可以选择的跳跃步数，farthest标记了所有可选择跳跃步数[i..end]中能够跳到的最远距离，jumps记录了跳跃次数。
*/

func jump(nums []int) int {
	n := len(nums)
	end := 0
	farthest := 0
	jumps := 0
	for i := 0; i < n-1; i++ {
		farthest = Max(nums[i]+i, farthest)
		if end == i {
			jumps++
			end = farthest
		}
	}
	return jumps
}

var memo []int

/**
请问你最少要跳多少次，才能跳过去？

我们先来说说动态规划的思路，采用自顶向下的递归动态规划，可以这样定义一个dp函数：

// 定义：从索引 p 跳到最后一格，至少需要 dp(nums, p) 步
int dp(vector<int>& nums, int p);
我们想求的结果就是dp(nums, 0)，base case 就是当p超过最后一格时，不需要跳跃：

if (p >= nums.size() - 1) {
    return 0;
}

*/
func jumpdp(nums []int) int {
	n := len(nums)
	// 备忘录都初始化为 n，相当于 INT_MAX
	// 因为从 0 调到 n - 1 最多 n - 1 步
	memo = make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = n
	}
	return dp45(nums, 0)
}

func dp45(nums []int, p int) int {
	n := len(nums)
	// base case
	if p >= n-1 {
		return 0
	}
	// 子问题已经计算过
	if nums[p] != n {
		return nums[p]
	}
	steps := nums[p]
	// 你可以选择跳 1 步，2 步...
	for i := 1; i <= steps; i++ {
		// 穷举每一个选择
		// 计算每一个子问题的结果
		subProblem := dp45(nums, p+i)
		// 取其中最小的作为最终结果
		memo[p] = Min(memo[p], subProblem+1)
	}
	return memo[p]
}
