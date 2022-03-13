package Code

import "testing"

/**
nums1中数字x的 下一个更大元素 是指x在nums2 中对应位置 右侧 的 第一个 比x大的元素。

给你两个 没有重复元素 的数组nums1 和nums2 ，下标从 0 开始计数，其中nums1是nums2的子集。

对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定 nums2[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。

返回一个长度为nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。



示例 1：

输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
输出：[-1,3,-1]
解释：nums1 中每个值的下一个更大元素如下所述：
- 4 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
- 1 ，用加粗斜体标识，nums2 = [1,3,4,2]。下一个更大元素是 3 。
- 2 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
示例 2：

输入：nums1 = [2,4], nums2 = [1,2,3,4].
输出：[3,-1]
解释：nums1 中每个值的下一个更大元素如下所述：
- 2 ，用加粗斜体标识，nums2 = [1,2,3,4]。下一个更大元素是 3 。
- 4 ，用加粗斜体标识，nums2 = [1,2,3,4]。不存在下一个更大元素，所以答案是 -1 。


提示：

1 <= nums1.length <= nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 104
nums1和nums2中所有整数 互不相同
nums1 中的所有整数同样出现在 nums2 中

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-greater-element-i
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test496(t *testing.T) {
	t.Log(nextGreaterNumber([]int{2, 1, 2, 4, 3}))
	t.Log(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	finalRes := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res := make([]int, len(nums2))
		stack := make([]int, 0, len(nums1))
		for j := len(nums2) - 1; j >= 0; j-- {
			for len(stack) > 0 && stack[len(stack)-1] <= nums2[j] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				res[j] = -1
			} else {
				res[j] = stack[len(stack)-1]
			}
			stack = append(stack, nums2[j])
			if nums1[i] == nums2[j] {
				finalRes[i] = res[j]
				break
			}
		}
	}
	return finalRes
}

func nextGreaterNumber(nums []int) []int {
	res := make([]int, len(nums))
	stack := make([]int, 0, len(nums))
	// 倒着往栈里放
	for j := len(nums) - 1; j >= 0; j-- {
		// 判定个子高矮
		for len(stack) > 0 && stack[len(stack)-1] <= nums[j] {
			// 矮个起开，反正也被挡着了。。。
			stack = stack[:len(stack)-1]
		}

		// nums[i] 身后的 next great number
		if len(stack) == 0 {
			res[j] = -1
		} else {
			res[j] = stack[len(stack)-1]
		}
		// 进队，接受之后的身高判定吧！
		stack = append(stack, nums[j])
	}
	return res
}
