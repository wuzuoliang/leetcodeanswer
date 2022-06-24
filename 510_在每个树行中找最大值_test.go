package Code

import (
	"math"
	"testing"
)

func Test510(t *testing.T) {
	tree := CreateFullTreeRoot([]int{1, 3, 2, 5, 3 - 1, 9})
	t.Log(largestValues(tree))
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	list := make([]*TreeNode, 0)
	list = append(list, root)
	results := make([]int, 0)
	for len(list) > 0 {
		lineTotal := len(list)
		lineMax := math.MinInt
		for i := 0; i < lineTotal; i++ {
			curNode := list[0]
			lineMax = Max(lineMax, curNode.Val)
			if curNode.Left != nil {
				list = append(list, curNode.Left)
			}
			if curNode.Right != nil {
				list = append(list, curNode.Right)
			}
			list = list[1:]
		}
		results = append(results, lineMax)
	}
	return results
}
