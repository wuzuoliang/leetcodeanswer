package Code

func buildTree105(preorder []int, inorder []int) *TreeNode {
	return build105(preorder, 0, len(preorder)-1,
		inorder, 0, len(inorder)-1)
}

func build105(preorder []int, preleft, preright int, inorder []int, inleft, inright int) *TreeNode {
	if preleft > preright {
		return nil
	}

	rootValue := preorder[preleft]
	rootIndex := -1
	for i := inleft; i <= inright; i++ {
		if inorder[i] == rootValue {
			rootIndex = i
			break
		}
	}
	leftSize := rootIndex - inleft
	root := &TreeNode{Val: rootValue}

	root.Left = build105(preorder, preleft+1, preleft+leftSize, inorder, inleft, inright-1)
	root.Right = build105(preorder, preleft+leftSize+1, preright, inorder, rootIndex+1, inright)

	return root
}
