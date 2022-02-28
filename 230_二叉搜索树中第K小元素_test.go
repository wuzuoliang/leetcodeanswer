package Code

import "testing"

func Test230(t *testing.T) {
	tree := CreateFullTreeRoot([]int{4, 1, 6, 2, 3, 5, 7})
	t.Log(kthSmallest(tree, 6))
}

var orderIndex int
var minK int

func kthSmallest(root *TreeNode, k int) int {
	orderIndex = 0
	minK = 0
	preOrder(root, k)
	return minK
}

func preOrder(root *TreeNode, k int) {
	if root == nil {
		return
	}
	preOrder(root.Left, k)
	orderIndex++
	if orderIndex == k {
		minK = root.Val
		return
	}
	preOrder(root.Right, k)
}
