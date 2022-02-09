package Code

/**
https://www.bilibili.com/video/BV1eg411w7gn?p=31
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
