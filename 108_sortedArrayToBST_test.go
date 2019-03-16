package Code

import "testing"

/**
Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example:

Given the sorted array: [-10,-3,0,5,9],

One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:

      0
     / \
   -3   9
   /   /
 -10  5
*/

func Test_sortedArrayToBST(t *testing.T) {

}
func sortedArrayToBST(nums []int) *TreeNode {
	return toBst(nums, 0, len(nums)-1)
}

func toBst(nums []int, left, right int) *TreeNode {
	if right < left {
		return nil
	}

	mid := (left + right) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  toBst(nums, left, mid-1),
		Right: toBst(nums, mid+1, right),
	}
}
