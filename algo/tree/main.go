package main

import (
	list "container/list"
	"fmt"
)

/**
树的定义：

有节点间的层次关系，分为父节点和子节点。
有唯一一个根节点，该根节点没有父节点。
除了根节点，每个节点有且只有一个父节点。
每一个节点本身以及它的后代也是一棵树，是一个递归的结构。
没有后代的节点称为叶子节点，没有节点的树称为空树。
二叉树：每个节点最多只有两个儿子节点的树。

满二叉树：叶子节点与叶子节点之间的高度差为 0 的二叉树，即整棵树是满的，树呈满三角形结构。在国外的定义，非叶子节点儿子都是满的树就是满二叉树。我们以国内为准。

完全二叉树：完全二叉树是由满二叉树而引出来的，设二叉树的深度为 k，除第 k 层外，其他各层的节点数都达到最大值，且第 k 层所有的节点都连续集中在最左边。


二叉树的数学特征
高度为 h≥0 的二叉树至少有 h+1 个结点，比如最不平衡的二叉树就是退化的线性链表结构，所有的节点都只有左儿子节点，或者所有的节点都只有右儿子节点。
高度为 h≥0 的二叉树至多有 2^h+1 个节点，比如这棵树是满二叉树。
含有 n≥1 个结点的二叉树的高度至多为 n-1，由 1 退化的线性链表可以反推。
含有 n≥1 个结点的二叉树的高度至少为 logn，由 2 满二叉树可以反推。
在二叉树的第 i 层，至多有 2^(i-1) 个节点，比如该层是满的。

对于一棵有 n 个节点的完全二叉树，从上到下，从左到右进行序号编号，对于任一个节点，编号 i=0 表示树根节点，编号 i 的节点的左右儿子节点编号分别为：2i+1,2i+2，父亲节点编号为：i/2，整除操作去掉小数。
*/

// 二叉树
type TreeNode struct {
	Data  string    // 节点用来存放数据
	Left  *TreeNode // 左子树
	Right *TreeNode // 右字树
}

func main() {
	t := &TreeNode{Data: "A"}
	t.Left = &TreeNode{Data: "B"}
	t.Right = &TreeNode{Data: "C"}
	t.Left.Left = &TreeNode{Data: "D"}
	t.Left.Right = &TreeNode{Data: "E"}
	t.Right.Left = &TreeNode{Data: "F"}
	PreOrder(t)
	fmt.Println("---")
	MidOrder(t)
	fmt.Println("---")
	PostOrder(t)
	fmt.Println("---")
	LayerOrder(t)
}

func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	fmt.Println(tree.Data)
	PreOrder(tree.Left)
	PreOrder(tree.Right)
}

func MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	PreOrder(tree.Left)
	fmt.Println(tree.Data)
	PreOrder(tree.Right)
}

func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	PreOrder(tree.Left)
	PreOrder(tree.Right)
	fmt.Println(tree.Data)
}

/**
层次遍历较复杂，用到一种名叫广度遍历的方法，需要使用辅助的先进先出的队列。

先将树的根节点放入队列。
从队列里面 remove 出节点，先打印节点值，如果该节点有左子树节点，左子树入栈，如果有右子树节点，右子树入栈。
重复2，直到队列里面没有元素。
*/
func LayerOrder(treeNode *TreeNode) {
	if treeNode == nil {
		return
	}
	q := list.New()
	q.PushBack(treeNode)
	for q.Len() > 0 {
		e := q.Front()
		node := e.Value.(*TreeNode)
		if node.Left != nil {
			q.PushBack(node.Left)
		}
		if node.Right != nil {
			q.PushBack(node.Right)
		}
		q.Remove(e)
		fmt.Println(node.Data)
	}
}
