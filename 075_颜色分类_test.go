package Code

import "testing"

/**
给定一个包含红色、白色和蓝色、共n 个元素的数组nums，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。

必须在不使用库的sort函数的情况下解决这个问题。

示例 1：

输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]
示例 2：

输入：nums = [2,0,1]
输出：[0,1,2]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-colors
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test75(t *testing.T) {
	sortColors([]int{2, 0, 2, 1, 1, 0})
}
func sortColors(nums []int) {
	size := len(nums)
	if size < 2 {
		return
	}

	// all in [0, zero) = 0
	// all in [zero, i) = 1
	// all in [two, len - 1] = 2

	zero := 0
	two := size
	i := 0
	for i < two {
		if nums[i] == 0 {
			swapNums(nums, zero, i)
			zero++
			i++
		} else if nums[i] == 1 {
			i++
		} else {
			two--
			swapNums(nums, i, two)
		}
	}
}
func swapNums(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
