package Code

import "testing"

/**
小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为root。

除了root之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。

给定二叉树的root。返回在不触动警报的情况下，小偷能够盗取的最高金额。



示例 1:



输入: root = [3,2,3,null,3,null,1]
输出: 7
解释:小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7
示例 2:



输入: root = [3,4,5,1,3,null,1]
输出: 9
解释:小偷一晚能够盗取的最高金额 4 + 5 = 9


提示：

树的节点数在[1, 104] 范围内
0 <= Node.val <= 104

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test337(t *testing.T) {
}

/**
Map<TreeNode, Integer> memo = new HashMap<>();
public int rob(TreeNode root) {
    if (root == null) return 0;
    // 利用备忘录消除重叠子问题
    if (memo.containsKey(root))
        return memo.get(root);
    // 抢，然后去下下家
    int do_it = root.val
        + (root.left == null ?
            0 : rob(root.left.left) + rob(root.left.right))
        + (root.right == null ?
            0 : rob(root.right.left) + rob(root.right.right));
    // 不抢，然后去下家
    int not_do = rob(root.left) + rob(root.right);

    int res = Math.max(do_it, not_do);
    memo.put(root, res);
    return res;
}
*/

func rob337(root *TreeNode) int {
	var memo map[*TreeNode]int
	memo = make(map[*TreeNode]int)
	return robTree(root, memo)

}

func robTree(root *TreeNode, memo map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if _, ok := memo[root]; ok {
		return memo[root]
	}

	do := root.Val
	if root.Left != nil {
		do += robTree(root.Left.Left, memo) + robTree(root.Left.Right, memo)
	}
	if root.Right != nil {
		do += robTree(root.Right.Left, memo) + robTree(root.Right.Right, memo)
	}

	notDo := robTree(root.Left, memo) + robTree(root.Right, memo)

	res := Max(do, notDo)
	memo[root] = res
	return res
}

/**
int rob(TreeNode root) {
    int[] res = dp(root);
    return Math.max(res[0], res[1]);
}

// 返回一个大小为 2 的数组 arr
// arr[0] 表示不抢 root 的话，得到的最大钱数
// arr[1] 表示抢 root 的话，得到的最大钱数
int[] dp(TreeNode root) {
if (root == null)
return new int[]{0, 0};
int[] left = dp(root.left);
int[] right = dp(root.right);
// 抢，下家就不能抢了
int rob = root.val + left[0] + right[0];
// 不抢，下家可抢可不抢，取决于收益大小
int not_rob = Math.max(left[0], left[1])
+ Math.max(right[0], right[1]);

return new int[]{not_rob, rob};
}
*/
