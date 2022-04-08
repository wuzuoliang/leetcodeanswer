package Code

/**
给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。

树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。



示例 1：



输入：root = [1,null,3,2,4,null,5,6]
输出：[[1],[3,2,4],[5,6]]
示例 2：



输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
输出：[[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]


提示：

树的高度不会超过1000
树的节点总数在 [0,10^4] 之间

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func levelOrder429(root *NTreeNode) [][]int {
	if root == nil {
		return nil
	}
	results := make([][]int, 0)
	q := make([]*NTreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		levelQ := len(q)
		innerLevel := make([]int, 0, levelQ)
		for i := 0; i < levelQ; i++ {
			q0 := q[0]
			q = q[1:]
			for _, v := range q0.Children {
				q = append(q, v)
			}
			innerLevel = append(innerLevel, q0.Val)
		}
		results = append(results, innerLevel)
	}
	return results
}
