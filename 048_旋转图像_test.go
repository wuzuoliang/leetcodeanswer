package Code

import (
	"testing"
)

/**
第48题：旋转图像
给定一个 n × n 的二维矩阵表示一个图像。
说明：

你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。


示例 1:

给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]

示例 2:

给定 matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],

原地旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
*/

func Test48(t *testing.T) {
	src := [][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}
	rotate(src)
	t.Log(src)
	rotateNi(src)
	t.Log(src)
}

// 顺时针90度
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		reverse48(matrix[i])
	}
}

// 逆时针90度
func rotateNi(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			matrix[i][j], matrix[n-j-1][n-i-1] = matrix[n-j-1][n-i-1], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		reverse48(matrix[i])
	}
}

func reverse48(arr []int) {
	i := 0
	j := len(arr) - 1

	for j > i {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
