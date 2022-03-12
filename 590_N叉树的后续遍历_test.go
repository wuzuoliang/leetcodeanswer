package Code

type NTreeNode struct {
	Val      int
	Children []*NTreeNode
}

func postorder(root *NTreeNode) []int {
	res := make([]int, 0)
	postorderNtree(root, &res)
	return res
}

func postorderNtree(root *NTreeNode, res *[]int) {
	if root == nil {
		return
	}
	for i := range root.Children {
		postorderNtree(root.Children[i], res)
	}
	*res = append(*res, root.Val)
}
