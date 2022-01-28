package Code

import (
	"testing"
)

/**
Given a non-empty array containing only positive integers, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

Note:
Both the array size and each of the array element will not exceed 100.

Example 1:

Input: [1, 5, 11, 5]

Output: true

Explanation: The array can be partitioned as [1, 5, 5] and [11].
Example 2:

Input: [1, 2, 3, 5]

Output: false

Explanation: The array cannot be partitioned into equal sum subsets.
*/

func TestCanPartition(t *testing.T) {
	t.Log(canPartition([]int{1, 2, 5, 2}))
	t.Log(canPartition([]int{1, 5, 11, 5}))
	t.Log(canPartition([]int{1, 2, 3, 5}))
	t.Log(canPartition2([]int{1, 8, 9, 10, 7, 3}))
}

/**
dp[i][j] = x表示，对于前i个物品，当前背包的容量为j时，若x为true，则说明可以恰好将背包装满，若x为false，则说明不能恰好将背包装满。

比如说，如果dp[4][9] = true，其含义为：对于容量为 9 的背包，若只是用前 4 个物品，可以有一种方法把背包恰好装满。

如果不把nums[i]算入子集，或者说你不把这第i个物品装入背包，那么是否能够恰好装满背包，取决于上一个状态dp[i-1][j]，继承之前的结果。

如果把nums[i]算入子集，或者说你把这第i个物品装入了背包，那么是否能够恰好装满背包，取决于状态dp[i - 1][j-nums[i-1]]。

*/
func canPartition(nums []int) bool {
	lens := len(nums)
	if len(nums) == 0 {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	/**
	\ 0 1 2 3 4 5
	0 y n n n n n
	1 y y n n n n
	2 y y y y n n
	5 y y y y n y
	2 y y y y y y
	*/
	sum = sum / 2
	dp := make([][]int, lens+1)
	for i := 0; i <= lens; i++ {
		dp[i] = make([]int, sum+1)
		dp[i][0] = 1
	}

	for i := 1; i <= lens; i++ {
		for j := 1; j <= sum; j++ {
			if j-nums[i-1] < 0 {
				// 背包容量不足，不能装入第 i 个物品
				dp[i][j] = dp[i-1][j]
			} else if j == nums[i-1] {
				dp[i][j] = 1
			} else {
				// 装入或不装入背包
				dp[i][j] = dp[i-1][j] | dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[lens-1][sum] == 1
}

// 注意到dp[i][j]都是通过上一行dp[i-1][..]转移过来的，之前的数据都不会再使用了。
//
// 所以，我们可以进行状态压缩，将二维dp数组压缩为一维，节约空间复杂度：
func canPartition2(nums []int) bool {
	lens := len(nums)
	if len(nums) == 0 {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	dp := make([]int, sum+1)
	dp[0] = 1
	for i := 0; i < lens; i++ {
		for j := sum; j >= 0; j-- {
			if j-nums[i] >= 0 {
				dp[j] = dp[j] | dp[j-nums[i]]
			}
		}
	}
	/**
	这就是状态压缩，其实这段代码和之前的解法思路完全相同，只在一行dp数组上操作，i每进行一轮迭代，dp[j]其实就相当于dp[i-1][j]，所以只需要一维数组就够用了。

	唯一需要注意的是j应该从后往前反向遍历，因为每个物品（或者说数字）只能用一次，以免之前的结果影响其他的结果。
	*/
	return dp[sum] == 1
}
