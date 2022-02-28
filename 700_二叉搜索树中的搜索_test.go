package Code

import "testing"

/**
第700题：二叉搜索树中的搜索
给定二叉搜索树（BST）的根节点和一个值。你需要在 BST 中找到节点值等于给定值的节点。返回以该节点为根的子树。如果节点不存在，则返回 NULL 。
示例：

给定二叉搜索树:

        4
       / \
      2   7
     / \
    1   3

和值: 2

你应该返回如下子树:

      2
     / \
    1   3
在上述示例中，如果要找的值是 5 ，但因为没有节点值为 5 ，我们应该返回 NULL 。
*/

func Test700(t *testing.T) {
	tree := CreateFullTreeRoot([]int{5, 2, 7, 1, 3, 6, 8})
	t.Log(isInBST(tree, 9))
}

func isInBST(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == target {
		return root
	}
	if root.Val < target {
		return isInBST(root.Right, target)
	} else {
		return isInBST(root.Left, target)
	}
}
