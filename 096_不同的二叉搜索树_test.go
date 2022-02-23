package Code

import "testing"

/**
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
*/

func Test96(t *testing.T) {
	t.Log(numTrees(3)) //5
	t.Log(numTrees(1)) //1
}
func numTrees(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]

	/**
	C := 1
	for i := 0; i < n; i++ {
		C = C * 2 * (2 * i + 1) / (i + 2);
	}
	return C

	作者：LeetCode-Solution
	链接：https://leetcode-cn.com/problems/unique-binary-search-trees/solution/bu-tong-de-er-cha-sou-suo-shu-by-leetcode-solution/
	来源：力扣（LeetCode）
	著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
	*/
}
