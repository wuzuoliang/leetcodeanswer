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

/**
举个例子，比如给算法输入n = 5，也就是说用{1,2,3,4,5}这些数字去构造 BST。

首先，这棵 BST 的根节点总共有几种情况？

显然有 5 种情况对吧，因为每个数字都可以作为根节点。

比如说我们固定3作为根节点，这个前提下能有几种不同的 BST 呢？

根据 BST 的特性，根节点的左子树都比根节点的值小，右子树的值都比根节点的值大。

所以如果固定3作为根节点，左子树节点就是{1,2}的组合，右子树就是{4,5}的组合。

左子树的组合数和右子树的组合数乘积就是3作为根节点时的 BST 个数。

我们这是说了3为根节点这一种特殊情况，其实其他的节点也是一样的。

那你可能会问，我们可以一眼看出{1,2}和{4,5}有几种组合，但是怎么让算法进行计算呢？

其实很简单，只需要递归就行了，我们可以写这样一个函数：

// 定义：闭区间 [lo, hi] 的数字能组成 count(lo, hi) 种 BST
int count(int lo, int hi);
根据这个函数的定义，结合刚才的分析，可以写出代码：


int numTrees(int n) {
// 计算闭区间 [1, n] 组成的 BST 个数
return count(1, n);
}

int count(int lo, int hi) {
// base case
if (lo > hi) return 1;

int res = 0;
for (int i = lo; i <= hi; i++) {
// i 的值作为根节点 root
int left = count(lo, i - 1);
int right = count(i + 1, hi);
// 左右子树的组合数乘积是 BST 的总数
res += left * right;
}

return res;
}
注意 base case，显然当lo > hi闭区间[lo, hi]肯定是个空区间，也就对应着空节点 null，虽然是空节点，但是也是一种情况，所以要返回 1 而不能返回 0。

这样，题目的要求已经实现了，但是时间复杂度非常高，肯定存在重叠子问题。

前文动态规划相关的问题多次讲过消除重叠子问题的方法，无非就是加一个备忘录：
// 备忘录
int[][] memo;

int numTrees(int n) {
    // 备忘录的值初始化为 0
    memo = new int[n + 1][n + 1];
    return count(1, n);
}

int count(int lo, int hi) {
    if (lo > hi) return 1;
    // 查备忘录
    if (memo[lo][hi] != 0) {
        return memo[lo][hi];
    }

    int res = 0;
    for (int mid = lo; mid <= hi; mid++) {
        int left = count(lo, mid - 1);
        int right = count(mid + 1, hi);
        res += left * right;
    }
    // 将结果存入备忘录
    memo[lo][hi] = res;

    return res;
}
*/
