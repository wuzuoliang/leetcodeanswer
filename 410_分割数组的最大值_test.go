package Code

import "testing"

/**
给定一个非负整数数组 nums 和一个整数m ，你需要将这个数组分成m个非空的连续子数组。

设计一个算法使得这m个子数组各自和的最大值最小。



示例 1：

输入：nums = [7,2,5,10,8], m = 2
输出：18
解释：
一共有四种方法将 nums 分割为 2 个子数组。
其中最好的方式是将其分为 [7,2,5] 和 [10,8] 。
因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。
示例 2：

输入：nums = [1,2,3,4,5], m = 2
输出：9
示例 3：

输入：nums = [1,4,4], m = 3
输出：4


提示：

1 <= nums.length <= 1000
0 <= nums[i] <= 106
1 <= m <= min(50, nums.length)

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/split-array-largest-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test410(t *testing.T) {
	t.Log(splitArray([]int{1, 4, 4}, 3))
}

func splitArray(nums []int, m int) int {
	min := 0
	max := 0
	for _, v := range nums {
		min = Max(min, v)
		max += v
	}

	for min < max {
		mid := min + (max-min)/2
		n := split(nums, mid)
		if n == m {
			max = mid
		} else if n < m {
			max = mid
		} else {
			min = mid + 1
		}
	}
	return min
}

/* 辅助函数，若限制最大子数组和为 max，
计算 nums 至少可以被分割成几个子数组 */
func split(nums []int, max int) int {
	// 至少可以分割的子数组数量
	count := 1
	// 记录每个子数组的元素和
	sum := 0
	for i := 0; i < len(nums); i++ {
		if sum+nums[i] > max {
			// 如果当前子数组和大于 max 限制
			// 则这个子数组不能再添加元素了
			count++
			sum = nums[i]
		} else {
			// 当前子数组和还没达到 max 限制
			// 还可以添加元素
			sum += nums[i]
		}
	}
	return count
}
