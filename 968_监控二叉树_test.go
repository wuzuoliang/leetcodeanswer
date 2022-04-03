package Code

import "math"

/**
给定一个二叉树，我们在树的节点上安装摄像头。

节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。

计算监控树的所有节点所需的最小摄像头数量。



示例 1：



输入：[0,0,null,0,0]
输出：1
解释：如图所示，一台摄像头足以监控所有节点。
示例 2：



输入：[0,0,null,0,null,0,null,null,0]
输出：2
解释：需要至少两个摄像头来监视树的所有节点。 上图显示了摄像头放置的有效位置之一。

提示：

给定树的节点数的范围是[1, 1000]。
每个节点的值都是 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-cameras
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
来看看这个状态应该如何转移，先来看看每个节点可能有几种状态：

有如下三种：

该节点无覆盖
本节点有摄像头
本节点有覆盖

我们分别有三个数字来表示：

0：该节点无覆盖
1：本节点有摄像头
2：本节点有覆盖

主要有如下四类情况：

情况1：左右节点都有覆盖
情况2：左右节点至少有一个无覆盖的情况
情况3：左右节点至少有一个有摄像头
情况4：头结点没有覆盖

*/
func minCameraCover(root *TreeNode) int {
	var dfs func(*TreeNode) (a, b, c int)
	dfs = func(node *TreeNode) (a, b, c int) {
		if node == nil {
			return math.MaxInt64 / 2, 0, 0
		}
		lefta, leftb, leftc := dfs(node.Left)
		righta, rightb, rightc := dfs(node.Right)
		a = leftc + rightc + 1
		b = Min(a, Min(lefta+rightb, righta+leftb))
		c = Min(a, leftb+rightb)
		return
	}
	_, ans, _ := dfs(root)
	return ans
}
