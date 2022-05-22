package Code

import "testing"

/**
给你一个二进制字符串数组 strs 和两个整数 m 和 n 。

请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。

如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。



示例 1：

输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
输出：4
解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 n 的值 3 。
示例 2：

输入：strs = ["10", "0", "1"], m = 1, n = 1
输出：2
解释：最大的子集是 {"0", "1"} ，所以答案是 2 。


提示：

1 <= strs.length <= 600
1 <= strs[i].length <= 100
strs[i]仅由'0' 和'1' 组成
1 <= m, n <= 100


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ones-and-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test474(t *testing.T) {
	t.Log(findMaxForm([]string{"10", "0", "1"}, 1, 1))
}

func findMaxForm(strs []string, m int, n int) int {
	// 定义数组
	// dp[i][j]：最多有i个0和j个1的strs的最大子集的大小为dp[i][j]。
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	// 遍历
	for i := 0; i < len(strs); i++ {
		zeroNum, oneNum := 0, 0
		//计算0,1 个数
		//或者直接strings.Count(strs[i],"0")
		for _, v := range strs[i] {
			if v == '0' {
				zeroNum++
			} else {
				oneNum++
			}
		}
		// 从后往前 遍历背包容量
		for j := m; j >= zeroNum; j-- {
			for k := n; k >= oneNum; k-- {
				// 推导公式
				dp[j][k] = Max(dp[j][k], dp[j-zeroNum][k-oneNum]+1)
			}
		}
	}
	return dp[m][n]
}
