package Code

import (
	"math"
	"testing"
)

/**
验证二叉搜索树
给定一个二叉树，判断其是否是一个有效的二叉搜索树。
示例 1:

输入:
   5
  / \
 1   4
    / \
   3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。

示例 2:

输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。

首先看完题目，我们很容易想到 遍历整棵树，比较所有节点，通过 左节点值<节点值，右节点值＞节点值 的方式来进行求解。但是这种解法是错误的，因为对于任意一个节点，我们不光需要左节点值小于该节点，并且左子树上的所有节点值都需要小于该节点。（右节点一致）所以我们在此引入上界与下界，用以保存之前的节点中出现的最大值与最小值。
*/

func TestIsBST(t *testing.T) {
	t.Log(isValidBST(CreateTreeRoot([]int{1, 2, 3, 4, 5, 6})))
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBST(root, math.MinInt64, math.MaxInt64)
}

func isBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if min >= root.Val || max <= root.Val {
		return false
	}
	return isBST(root.Left, min, root.Val) && isBST(root.Right, root.Val, max)
}
