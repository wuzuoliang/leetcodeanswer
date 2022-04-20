package Code

import "testing"

/**
编写一个高效的算法来搜索mxn矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-a-2d-matrix-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test240(t *testing.T) {
	t.Log(searchMatrix([][]int{{1, 4, 7, 11}, {2, 5, 8, 12}, {3, 6, 9, 16}, {10, 13, 14, 17}}, 5)) // true
	t.Log(searchMatrix([][]int{{5}, {6}}, 6))                                                      // true
	t.Log(searchMatrix([][]int{{-5}}, -10))                                                        // true
	t.Log(searchMatrix([][]int{
		{3, 3, 8, 13, 13, 18},
		{4, 5, 11, 13, 18, 20},
		{9, 9, 14, 15, 23, 23},
		{13, 18, 22, 22, 25, 27},
		{18, 22, 23, 28, 30, 33},
		{21, 25, 28, 30, 35, 35},
		{24, 25, 33, 36, 37, 40}}, 21))
}

// https://www.bilibili.com/video/BV1XU4y1d7E1?p=41
// 升级问题，返回最小的K个的数
func searchMatrix(matrix [][]int, target int) bool {
	row := 0
	line := len(matrix)
	col := len(matrix[0]) - 1
	for row < line && col >= 0 {
		if target > matrix[row][col] {
			row++
		} else if target < matrix[row][col] {
			col--
		} else {
			return true
		}
	}
	return false
}
