package Code

import "testing"

/**
给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的 下一个更大元素 。

数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1 。



示例 1:

输入: nums = [1,2,1]
输出: [2,-1,2]
解释: 第一个 1 的下一个更大的数是 2；
数字 2 找不到下一个更大的数；
第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
示例 2:

输入: nums = [1,2,3,4,3]
输出: [2,3,4,-1,4]

提示:

1 <= nums.length <= 104
-109 <= nums[i] <= 109

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-greater-element-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test503(t *testing.T) {
	t.Log(nextGreaterElements([]int{1, 2, 3, 4, 5}))
	t.Log(nextGreaterElements([]int{2, 1, 2, 4, 3}))

}

/*
   // 但是注意到只遍历一次序列是不够的，例如序列 [2,3,1][2,3,1]，最后单调栈中将剩余 [3,1][3,1]，其中元素 [1][1] 的下一个更大元素还是不知道的。
   所以 循环遍历2n-1个，
   如果栈为空，则把当前元素放入栈内；
   如果栈不为空，则需要判断当前元素和栈顶元素的大小：
   如果当前元素比栈顶大，说明当前元素是前面一些元素的「下一个更大元素」，则逐个弹出栈顶元素，直到当前元素比栈顶元素小为止。
   如果当前元素比栈顶元素小：说明当前元素的「下一个更大元素」与栈顶元素相同，则把当前元素入栈。
*/
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	stack := make([]int, 0, n)
	// 倒着往栈里放
	for j := 2*n - 1; j >= 0; j-- {
		// 判定个子高矮
		for len(stack) > 0 && stack[len(stack)-1] <= nums[j%n] {
			// 矮个起开，反正也被挡着了。。。
			stack = stack[:len(stack)-1]
		}

		// nums[i] 身后的 next great number
		if len(stack) == 0 {
			res[j%n] = -1
		} else {
			res[j%n] = stack[len(stack)-1]
		}
		// 进队，接受之后的身高判定吧！
		stack = append(stack, nums[j%n])
	}
	return res
}
