package Code

import "testing"

/**
第54题：螺旋矩阵
定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]

示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9, 10, 11, 12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]
*/

func TestSpiralMatric(t *testing.T) {
	t.Log(spiralMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	t.Log(spiralMatrix([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}

func spiralMatrix(s [][]int) []int {
	m := len(s)
	if m == 0 {
		return nil
	}
	n := len(s[0])
	results := make([]int, 0, m*n)
	head := 0
	botom := m - 1
	left := 0
	right := n - 1
	direct := 0
here:
	for {
		if len(results) == m*n {
			break
		}
		switch direct {
		case 0:
			// 右
			i := head
			for j := left; j <= right; j++ {
				results = append(results, s[i][j])
				if j == right {
					direct = (direct + 1) % 4
					head++
					goto here
				}
			}
		case 1:
			// 下
			j := right
			for i := head; i <= botom; i++ {
				results = append(results, s[i][j])
				if i == botom {
					direct = (direct + 1) % 4
					right--
					goto here
				}
			}
		case 2:
			// 左
			i := botom
			for j := right; j >= left; j-- {
				results = append(results, s[i][j])
				if j == left {
					direct = (direct + 1) % 4
					botom--
					goto here
				}
			}
		default:
			// 上
			j := left
			for i := botom; i >= head; i-- {
				results = append(results, s[i][j])
				if i == head {
					direct = (direct + 1) % 4
					left++
					goto here
				}
			}
		}
	}
	return results
}
