package Code

func inorderTraversal(root *TreeNode) []int {
	list := make([]int, 0)
	if root == nil {
		return list
	}

	first(root.Left, &list)
	list = append(list, root.Val)
	first(root.Right, &list)
	return list
}

func first(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	first(root.Left, list)
	*list = append(*list, root.Val)
	first(root.Right, list)
}
