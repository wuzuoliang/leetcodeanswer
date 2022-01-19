package Code

import "testing"

/**
给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。



示例 1：

输入：nums = [1,2,3,1], k = 3
输出：true
示例 2：

输入：nums = [1,0,1,1], k = 1
输出：true
示例 3：

输入：nums = [1,2,3,1,2,3], k = 2
输出：false




提示：

1 <= nums.length <= 105
-109 <= nums[i] <= 109
0 <= k <= 105


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/contains-duplicate-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestContainsNearbyDuplicate(t *testing.T) {
	t.Log(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	t.Log(containsNearbyDuplicate3([]int{1, 0, 1, 1}, 1))
	t.Log(containsNearbyDuplicate2([]int{1, 2, 3, 1, 2, 3}, 2))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}

	length := len(nums) - 1
	for left := 0; left < length; left++ {
		minRight := Min(left+k, length)
		for right := left + 1; right <= minRight; right++ {
			if nums[left] == nums[right] && right-left <= k {
				return true
			}
		}
	}
	return false
}

/**
哈希表
复杂度分析

时间复杂度：O(n)O(n)，其中 nn 是数组 \textit{nums}nums 的长度。需要遍历数组一次，对于每个元素，哈希表的操作时间都是 O(1)O(1)。

空间复杂度：O(n)O(n)，其中 nn 是数组 \textit{nums}nums 的长度。需要使用哈希表记录每个元素的最大下标，哈希表中的元素个数不会超过 nn。
*/
func containsNearbyDuplicate2(nums []int, k int) bool {
	existIndexMap := make(map[int]int, len(nums))

	for idx, num := range nums {
		if existIndex, ok := existIndexMap[num]; ok && idx-existIndex <= k {
			return true
		} else {
			existIndexMap[num] = idx
		}
	}

	return false

}

/**
滑动窗口
复杂度分析

时间复杂度：O(n)O(n)，其中 nn 是数组 \textit{nums}nums 的长度。需要遍历数组一次，对于每个元素，哈希集合的操作时间都是 O(1)。

空间复杂度：O(k)O(k)，其中 kk 是判断重复元素时允许的下标差的绝对值的最大值。需要使用哈希集合存储滑动窗口中的元素，任意时刻滑动窗口中的元素个数最多为 k + 1k+1 个。
*/
func containsNearbyDuplicate3(nums []int, k int) bool {
	set := map[int]struct{}{}
	for i, num := range nums {
		if i > k {
			delete(set, nums[i-k-1])
		}
		if _, ok := set[num]; ok {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}
