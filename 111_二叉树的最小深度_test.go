package Code

import (
	"container/list"
	"math"
)

/**
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明：叶子节点是指没有子节点的节点。



示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：2
示例 2：

输入：root = [2,null,3,null,4,null,5,null,6]
输出：5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-depth-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
1、为什么 BFS 可以找到最短距离，DFS 不行吗？

首先，你看 BFS 的逻辑，depth每增加一次，队列中的所有节点都向前迈一步，这保证了第一次到达终点的时候，走的步数是最少的。

DFS 不能找最短路径吗？其实也是可以的，但是时间复杂度相对高很多。

你想啊，DFS 实际上是靠递归的堆栈记录走过的路径，你要找到最短路径，肯定得把二叉树中所有树杈都探索完才能对比出最短的路径有多长对不对？

而 BFS 借助队列做到一次一步「齐头并进」，是可以在不遍历完整棵树的条件下找到最短距离的。

形象点说，DFS 是线，BFS 是面；DFS 是单打独斗，BFS 是集体行动。这个应该比较容易理解吧。

2、既然 BFS 那么好，为啥 DFS 还要存在？

BFS 可以找到最短距离，但是空间复杂度高，而 DFS 的空间复杂度较低。

还是拿刚才我们处理二叉树问题的例子，假设给你的这个二叉树是满二叉树，节点总数为N，对于 DFS 算法来说，空间复杂度无非就是递归堆栈，最坏情况下顶多就是树的高度，也就是O(logN)。

但是你想想 BFS 算法，队列中每次都会储存着二叉树一层的节点，这样的话最坏情况下空间复杂度应该是树的最底层节点的数量，也就是N/2，用 Big O 表示的话也就是O(N)。

由此观之，BFS 还是有代价的，一般来说在找最短路径的时候使用 BFS，其他时候还是 DFS 使用得多一些（主要是递归代码好写）。
*/

// DFS
func minDepthReverse(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = Min(minDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = Min(minDepth(root.Right), minD)
	}
	return minD + 1
}

// BFS
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lt := list.New()
	lt.PushBack(root)
	// root 本身就是一层，depth 初始化为 1

	depth := 1
	for lt.Len() != 0 {
		fElm := lt.Front()
		flen := lt.Len()
		/* 将当前队列中的所有节点向四周扩散 */

		for i := 0; i < flen; i++ {
			/* 判断是否到达终点 */

			if fElm.Value.(*TreeNode).Left == nil && fElm.Value.(*TreeNode).Right == nil {
				return depth
			}
			/* 将 cur 的相邻节点加入队列 */

			if fElm.Value.(*TreeNode).Left != nil {
				lt.PushBack(fElm.Value.(*TreeNode).Left)
			}
			if fElm.Value.(*TreeNode).Right != nil {
				lt.PushBack(fElm.Value.(*TreeNode).Right)
			}
			lt.Remove(fElm)
		}
		depth++
	}
	return depth
}
