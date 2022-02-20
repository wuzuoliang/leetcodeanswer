package Code

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBFS(t *testing.T) {
	convey.Convey("TestBFS", t, func() {
		t.Log(levelOrder(CreateTreeRoot([]int{1, 2, 3, 4, 5, 6, 7, 8, 5, 3})))
	})
}

/**
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。


示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]
示例 2：

输入：root = [1]
输出：[[1]]
示例 3：

输入：root = []
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	results := make([][]int, 0)
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		levelQ := len(q)
		innerLevel := make([]int, 0, levelQ)
		for i := 0; i < levelQ; i++ {
			q0 := q[0]
			q = q[1:]
			if q0.Left != nil {
				q = append(q, q0.Left)
			}
			if q0.Right != nil {
				q = append(q, q0.Right)
			}
			innerLevel = append(innerLevel, q0.Val)
		}
		results = append(results, innerLevel)
	}
	return results
}
