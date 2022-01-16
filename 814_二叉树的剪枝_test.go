package Code

/**
第814题：二叉树的剪枝
给定二叉树根结点 root ，此外树的每个结点的值要么是 0，要么是 1。返回移除了所有不包含 1 的子树的原二叉树。( 节点 X 的子树为 X 本身，以及所有 X 的后代。)
示例1:

输入: [1,null,0,0,1]
输出: [1,null,0,null,1]

说明:

给定的二叉树最多有 100 个节点。
每个节点的值只会为 0 或 1 。
*/

func pruneTree(root *TreeNode) *TreeNode {
	return deal(root)
}

func deal(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	node.Left = deal(node.Left)
	node.Right = deal(node.Right)
	if node.Left == nil && node.Right == nil && node.Val == 0 {
		return nil
	}
	return node
}
