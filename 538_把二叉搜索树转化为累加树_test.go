package Code

import "testing"

func Test538(t *testing.T) {
	tree := convertBST(CreateFullTreeRoot([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	t.Log(tree)
}

func travel(root *TreeNode, r int) int {
	if root == nil {
		return 0
	}

	right := travel(root.Right, r)
	tmp := root.Val
	root.Val += right + r
	left := travel(root.Left, root.Val)
	return left + right + tmp

}
func convertBST(root *TreeNode) *TreeNode {
	travel(root, 0)
	return root
}
