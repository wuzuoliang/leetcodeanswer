package Code

import "testing"

func Test538(t *testing.T) {
	tree := convertBST(CreateFullTreeRoot([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	t.Log(tree)
}

var sum int

func convertBST(root *TreeNode) *TreeNode {
	reverse538(root)
	return root
}

func reverse538(root *TreeNode) {
	if root == nil {
		return
	}
	reverse538(root.Right)
	sum += root.Val
	root.Val = sum
	reverse538(root.Left)
}
