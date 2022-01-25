package Code

/**
给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectTwo(root.Left, root.Right)
	return root
}

func connectTwo(left, right *Node) {
	if left == nil || right == nil {
		return
	}

	left.Next = right
	connectTwo(left.Left, left.Right)
	connectTwo(right.Left, right.Right)

	connectTwo(left.Right, right.Left)
}
