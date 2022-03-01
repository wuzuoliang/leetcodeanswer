package Code

import "testing"

/**
Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its bottom-up level order traversal as:
[
  [15,7],
  [9,20],
  [3]
]

*/

func Test_levelOrderBottom(t *testing.T) {

}
func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		orderLevel := make([]int, 0)
		orderLens := len(q)
		for i := 0; i < orderLens; i++ {
			t := q[0]
			orderLevel = append(orderLevel, t.Val)
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
			q = q[1:]
		}
		res = append(res, orderLevel)
	}

	i := 0
	j := len(res) - 1

	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--

	}
	return res
}
