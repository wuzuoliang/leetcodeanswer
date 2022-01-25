package Code

func buildTree106(inorder []int, postorder []int) *TreeNode {
	return build106(inorder, 0, len(inorder)-1,
		postorder, 0, len(postorder)-1)
}

func build106(inorder []int, inleft, inright int, postorder []int, pleft, pright int) *TreeNode {
	if inleft > inright {
		return nil
	}

	rootValue := postorder[pright]
	rootIndex := -1
	for i := inleft; i <= inright; i++ {
		if inorder[i] == rootValue {
			rootIndex = i
			break
		}
	}
	leftSize := rootIndex - inleft
	root := &TreeNode{Val: rootValue}

	root.Left = build106(inorder, inleft, rootIndex-1, postorder, pleft, pleft+leftSize-1)
	root.Right = build106(inorder, rootIndex+1, inright, postorder, pleft+leftSize, pright-1)

	return root
}
