package Code

import "testing"

func Test701(t *testing.T) {
	tree := CreateFullTreeRoot([]int{5, 2, 7, 1, 3, 6, 8})
	t.Log(isInBST(tree, 9))
	t.Log(isInBST(insertIntoBST(tree, 9), 9))
}
func insertIntoBST(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: value}
	}
	if root.Val <= value {
		root.Right = insertIntoBST(root.Right, value)
	} else {
		root.Left = insertIntoBST(root.Left, value)
	}
	return root
}
