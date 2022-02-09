package Code

func depthOfTree(root *TreeNode, maxDiameter *int) int {
	if root == nil {
		return 0
	}
	lDepth := depthOfTree(root.Left, maxDiameter)
	rDepth := depthOfTree(root.Right, maxDiameter)
	diameter := lDepth + rDepth

	if diameter > (*maxDiameter) {
		*maxDiameter = diameter
	}

	if lDepth > 0 && lDepth >= rDepth {
		return lDepth + 1
	}
	if rDepth > 0 && rDepth > lDepth {
		return rDepth + 1
	}
	return 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	var maxDiameter int
	depthOfTree(root, &maxDiameter)
	return maxDiameter
}
