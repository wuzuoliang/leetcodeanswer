package Code

import (
	"testing"
)

/**
输入为一个n * n的二维数组matrix，请你计算从第一行落到最后一行，经过的路径和最小为多少。
每次下降，可以向下、向左下、向右下三个方向移动一格。也就是说，可以从matrix[i][j]降到matrix[i+1][j]或matrix[i+1][j-1]或matrix[i+1][j+1]三个位置。
*/

func TestMinFallingPathSum(t *testing.T) {
	t.Log(minFallingPathSum([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}))
}

const max = 999999
const flag = 666666

var dp931 [][]int

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	res := max
	dp931 = make([][]int, n)
	for i := 0; i < n; i++ {
		dp931[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == n-1 {
				dp931[i][j] = matrix[i][j]
			}
		}
	}
	for j := 0; j < n; j++ {
		res = Min(res, getdp931(matrix, n-1, j))
	}
	return res
}

func getdp931(martix [][]int, i int, j int) int {
	n := len(martix)
	// 1、索引合法性检查
	if i < 0 || j < 0 || i >= n || j >= n {
		return max
	}
	// 2、base case
	if i == 0 {
		return dp931[0][j]
	}
	// 3、查找备忘录，防止重复计算
	if dp931[i][j] != flag {
		return dp931[i][j]
	}
	// 进行状态转移
	dp931[i][j] = dp931[i][j] + ThirdMin(getdp931(martix, i-1, j), getdp931(martix, i-1, j-1), getdp931(martix, i-1, j+1))

	return dp931[i][j]
}
