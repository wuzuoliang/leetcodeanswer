package Code

import "math"

/**
124. 二叉树中的最大路径和
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。
同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。
Example 1:

Input: [1,2,3]

       1
      / \
     2   3

Output: 6
Example 2:

Input: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

Output: 42
*/
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := math.MinInt32
	getPathSum(root, &max)
	return max
}

func getPathSum(root *TreeNode, maxSum *int) int {
	if root == nil {
		return math.MinInt32
	}
	left := getPathSum(root.Left, maxSum)
	right := getPathSum(root.Right, maxSum)

	currMax := Max(Max(left+root.Val, right+root.Val), root.Val)
	*maxSum = Max(*maxSum, Max(currMax, left+right+root.Val))
	return currMax
}
