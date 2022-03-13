package Code

import "testing"

/**
https://www.bilibili.com/video/BV1eg411w7gn?p=9
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意，必须在不复制数组的情况下原地对数组进行操作。



示例 1:

输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
示例 2:

输入: nums = [0]
输出: [0]


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/move-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test283(t *testing.T) {
	list := []int{1, 0, 2, 0, 3, 0, 4, 5, 0, 6}
	moveZeroes(list)
	t.Log(list)
}

func moveZeroes(nums []int) {
	left := 0
	right := 0
	n := len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
