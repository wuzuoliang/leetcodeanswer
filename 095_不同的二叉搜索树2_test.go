package Code

import "testing"

func Test95(t *testing.T) {
	t.Log(len(generateTrees(3))) //5
	t.Log(len(generateTrees(5))) //42
}

/**
// 主函数
public List<TreeNode> generateTrees(int n) {
if (n == 0) return new LinkedList<>();
// 构造闭区间 [1, n] 组成的 BST
return build(1, n);
}

// 构造闭区间 [lo, hi] 组成的 BST
List<TreeNode> build(int lo, int hi) {
List<TreeNode> res = new LinkedList<>();
// base case
if (lo > hi) {
res.add(null);
return res;
}

// 1、穷举 root 节点的所有可能。
for (int i = lo; i <= hi; i++) {
// 2、递归构造出左右子树的所有合法 BST。
List<TreeNode> leftTree = build(lo, i - 1);
List<TreeNode> rightTree = build(i + 1, hi);
// 3、给 root 节点穷举所有左右子树的组合。
for (TreeNode left : leftTree) {
for (TreeNode right : rightTree) {
// i 作为根节点 root 的值
TreeNode root = new TreeNode(i);
root.left = left;
root.right = right;
res.add(root);
}
}
}

return res;
}
*/
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return build95(1, n)
}
func build95(l, r int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if l > r {
		res = append(res, nil)
		return res
	}
	for i := l; i <= r; i++ {
		left := build95(l, i-1)
		right := build95(i+1, r)

		for _, lv := range left {
			for _, rv := range right {
				newTree := &TreeNode{Val: i, Left: lv, Right: rv}
				res = append(res, newTree)
			}
		}
	}
	return res
}
