package Code

import (
	"fmt"
	"testing"
)

var results [][]int

func TestTreeRoad(t *testing.T) {
	tree := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
			Left: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 8,
			Right: &TreeNode{
				Val: 9,
			},
			Left: &TreeNode{
				Val: 7,
			},
		},
	}
	treeRoad(tree, []int{})
	for _, v := range results {
		fmt.Println(v)
	}
}

func treeRoad(root *TreeNode, tmp []int) {
	if root.Left == nil && root.Right == nil {
		tmp = append(tmp, root.Val)
		newTmp := make([]int, len(tmp))
		copy(newTmp, tmp)
		results = append(results, newTmp)
		return
	}
	tmp = append(tmp, root.Val)
	if root.Left != nil {
		treeRoad(root.Left, tmp)
	}
	if root.Right != nil {
		treeRoad(root.Right, tmp)
	}
}
