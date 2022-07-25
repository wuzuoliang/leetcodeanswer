package Code

import (
	"testing"
)

/**
完全二叉树 是每一层（除最后一层外）都是完全填充（即，节点数达到最大）的，并且所有的节点都尽可能地集中在左侧。

设计一种算法，将一个新节点插入到一个完整的二叉树中，并在插入后保持其完整。

实现 CBTInserter 类:

CBTInserter(TreeNode root)使用头节点为root的给定树初始化该数据结构；
CBTInserter.insert(int v) 向树中插入一个值为Node.val == val的新节点TreeNode。使树保持完全二叉树的状态，并返回插入节点TreeNode的父节点的值；
CBTInserter.get_root() 将返回树的头节点。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/complete-binary-tree-inserter
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test919(t *testing.T) {
	ct := Constructor919(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}})
	//t.Log(ct.Insert(2))
	t.Log(ct.Insert(3))
	t.Log(ct.Insert(4))
	//t.Log(ct.Insert(5))
	//t.Log(ct.Insert(6))
	ct.Get_root()
}

type CBTInserter struct {
	root      *TreeNode
	candidate []*TreeNode
}

func Constructor919(root *TreeNode) CBTInserter {
	q := []*TreeNode{root}
	candidate := []*TreeNode{}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left == nil || node.Right == nil {
			candidate = append(candidate, node)
		}
	}
	return CBTInserter{root, candidate}
}

func (c *CBTInserter) Insert(val int) int {
	child := &TreeNode{Val: val}
	node := c.candidate[0]
	if node.Left == nil {
		node.Left = child
	} else {
		node.Right = child
		c.candidate = c.candidate[1:]
	}
	c.candidate = append(c.candidate, child)
	return node.Val
}

func (c *CBTInserter) Get_root() *TreeNode {
	return c.root
}
