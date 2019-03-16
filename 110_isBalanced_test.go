package Code

import "testing"

/**
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example 1:

Given the following tree [3,9,20,null,null,15,7]:

    3
   / \
  9  20
    /  \
   15   7
Return true.

Example 2:

Given the following tree [1,2,2,3,3,null,null,4,4]:

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
Return false.
*/

func Test_isBalanced(t *testing.T) {

}

func isBalanced(root *TreeNode) bool {
	isValid, _ := check(root)
	return isValid
}

func check(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	leftValid, leftDepth := check(root.Left)
	if !leftValid {
		return false, leftDepth + 1
	}
	rightValid, rightDepth := check(root.Right)
	if !rightValid {
		return false, rightDepth + 1
	}
	depth := leftDepth
	if leftDepth < rightDepth {
		depth = rightDepth
	}
	if diff := leftDepth - rightDepth; diff > 1 || diff < -1 {
		return false, depth
	}
	return true, depth + 1
}
