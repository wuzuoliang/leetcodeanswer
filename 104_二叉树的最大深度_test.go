package Code

import "testing"

/**
https://www.bilibili.com/video/BV1eg411w7gn?p=29
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its depth = 3.
*/

func Test_maxDepth(t *testing.T) {

}
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	return 1 + Max(MaxDepth(root.Left), MaxDepth(root.Right))
}
