package Code

import (
	"testing"
)

/**
给你一个二叉树的根结点root，请返回出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的子树元素和（不限顺序）。

一个结点的「子树元素和」定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。



示例 1：



输入: root = [5,2,-3]
输出: [2,-3,4]
示例2：



输入: root = [5,2,-5]
输出: [2]


提示:

节点数在[1, 104]范围内
-105<= Node.val <= 105

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/most-frequent-subtree-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test508(t *testing.T) {
	t1 := &TreeNode{Val: 5, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: -3}}
	t.Log(findFrequentTreeSum(t1))

	t2 := &TreeNode{Val: 5, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: -5}}
	t.Log(findFrequentTreeSum(t2))
}

func findFrequentTreeSum(root *TreeNode) (ans []int) {
	cnt := map[int]int{}
	maxCnt := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := node.Val + dfs(node.Left) + dfs(node.Right)
		cnt[sum]++
		if cnt[sum] > maxCnt {
			maxCnt = cnt[sum]
		}
		return sum
	}
	dfs(root)

	for s, c := range cnt {
		if c == maxCnt {
			ans = append(ans, s)
		}
	}
	return
}
