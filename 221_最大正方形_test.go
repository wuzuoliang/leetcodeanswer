package Code

import "testing"

/**
在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。



示例 1：


输入：matrix = [['1','0','1','0','0'],['1','0','1','1','1'],['1','1','1','1','1'],['1','0','0','1','0']]
输出：4
示例 2：


输入：matrix = [['0','1'],['1','0']]
输出：1
示例 3：

输入：matrix = [['0']]
输出：0


提示：

m == matrix.length
n == matrix[i].length
1 <= m, n <= 300
matrix[i][j] 为 '0' 或 '1'

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximal-square
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test221(t *testing.T) {
	t.Log(maximalSquare([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	})) //4

	t.Log(maximalSquare([][]byte{
		{'0', '1'},
		{'1', '0'},
	})) //1

	t.Log(maximalSquare([][]byte{
		{'0'},
	})) //0
}

/**
动态规划最重要的就是想清楚 d[i,j] 代表着什么，这个太重要了。 这道题中的 d[i, j] 代表的是，以坐标点(i,j) 为右下角的最大正方形，如果说你的 d[i,j] 是想 这个区域中 最大的正方形，那可就麻烦太多了。
如果该位置的值是 00，则 \textit{dp}(i, j) = 0dp(i,j)=0，因为当前位置不可能在由 11 组成的正方形中；

如果该位置的值是 11，则 \textit{dp}(i, j)dp(i,j) 的值由其上方、左方和左上方的三个相邻位置的 \textit{dp}dp 值决定。具体而言，当前位置的元素值等于三个相邻位置的元素中的最小值加 11，状态转移方程如下：

dp(i, j)=min(dp(i−1, j), dp(i−1, j−1), dp(i, j−1))+1
dp(i,j)=min(dp(i−1,j),dp(i−1,j−1),dp(i,j−1))+1

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/maximal-square/solution/zui-da-zheng-fang-xing-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	maxSide := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = Min(Min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}
	return maxSide * maxSide
}
