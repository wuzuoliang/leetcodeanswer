package Code

import "math"

/**
给定一个不重复的整数数组nums 。最大二叉树可以用下面的算法从nums 递归地构建:

创建一个根节点，其值为nums 中的最大值。
递归地在最大值左边的子数组前缀上构建左子树。
递归地在最大值 右边 的子数组后缀上构建右子树。
返回nums 构建的 最大二叉树 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return buildMaximumBinaryTree(nums, 0, len(nums)-1)
}
func buildMaximumBinaryTree(nums []int, l, h int) *TreeNode {
	if l > h {
		return nil
	}
	max := math.MinInt
	maxIndex := -1
	for i := l; i <= h; i++ {
		if nums[i] > max {
			max = nums[i]
			maxIndex = i
		}
	}

	return &TreeNode{
		Val:   max,
		Left:  buildMaximumBinaryTree(nums, l, maxIndex-1),
		Right: buildMaximumBinaryTree(nums, maxIndex+1, h),
	}

}
