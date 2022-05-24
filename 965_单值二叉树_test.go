package Code

/**
如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。

只有给定的树是单值二叉树时，才返回 true；否则返回 false。
*/
func isUnivalTree(root *TreeNode) bool {
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if root.Left != nil && root.Val != root.Left.Val {
			return false
		}
		if root.Right != nil && root.Val != root.Right.Val {
			return false
		}
		return dfs(root.Left) && dfs(root.Right)
	}
	return dfs(root)
}
