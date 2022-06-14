package Code

import "testing"

/**
给定一个m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。



示例 1：


输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
输出：[[1,0,1],[0,0,0],[1,0,1]]
示例 2：


输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]


提示：

m == matrix.length
n == matrix[0].length
1 <= m, n <= 200
-231 <= matrix[i][j] <= 231 - 1


进阶：

一个直观的解决方案是使用 O(mn)的额外空间，但这并不是一个好的解决方案。
一个简单的改进方案是使用 O(m+n) 的额外空间，但这仍然不是最好的解决方案。
你能想出一个仅使用常量空间的解决方案吗？

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/set-matrix-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test73(t *testing.T) {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	for _, v := range matrix {
		t.Log(v)
	}
}

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	zeros := make([][]int, 0)
	for i, l := range matrix {
		for y, c := range l {
			if c == 0 {
				zeros = append(zeros, []int{i, y})
			}
		}
	}

	for _, v := range zeros {
		x, y := v[0], v[1]
		for i := 0; i < m; i++ {
			matrix[i][y] = 0
		}
		for i := 0; i < n; i++ {
			matrix[x][i] = 0
		}
	}
}
