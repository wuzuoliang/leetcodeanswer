package Code

import (
	"testing"
)

/**
给你一个整数数组 nums 和一个整数k ，请你统计并返回该数组中和为k的连续子数组的个数。

示例 1：

输入：nums = [1,1,1], k = 2
输出：2
示例 2：

输入：nums = [1,2,3], k = 3
输出：2


提示：

1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subarray-sum-equals-k
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test560(t *testing.T) {
	t.Log(subarraySum([]int{1, 1, 1}, 2)) //2
	t.Log(subarraySum([]int{1, 2, 3}, 3)) //2
	t.Log(subarraySum([]int{1}, 0))       //0

}

/**
每个元素对应一个“前缀和”
遍历数组，根据当前“前缀和”，在 map 中寻找「与之相减 == k」的历史前缀和
当前“前缀和”与历史前缀和，差分出一个子数组，该历史前缀和出现过 c 次，就表示当前项找到 c 个子数组求和等于 k。
遍历过程中，c 不断加给 count，最后返回 count

作者：xiao_ben_zhu
链接：https://leetcode-cn.com/problems/subarray-sum-equals-k/solution/dai-ni-da-tong-qian-zhui-he-cong-zui-ben-fang-fa-y/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func subarraySum(nums []int, k int) int {
	count := 0
	hash := map[int]int{0: 1}
	preSum := 0

	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		if hash[preSum-k] > 0 {
			count += hash[preSum-k]
		}
		hash[preSum]++
	}
	return count
	/**
	其实我们不关心具体是哪两项的前缀和之差等于k，只关心等于 k 的前缀和之差出现的次数c，就知道了有c个子数组求和等于k
	每个元素对应一个“前缀和”
	遍历数组，根据当前“前缀和”，在 map 中寻找「与之相减 == k」的历史前缀和
	当前“前缀和”与历史前缀和，差分出一个子数组，该历史前缀和出现过 c 次，就表示当前项找到 c 个子数组求和等于 k。
	遍历过程中，c 不断加给 count，最后返回 count
	*/
}
