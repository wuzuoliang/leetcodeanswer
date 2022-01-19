package Code

import "testing"

/**
第59题：螺旋矩阵Ⅱ
给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例：

输入: 3
输出: [ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]
*/

func TestSpiralMatrix2(t *testing.T) {
	t.Log(spiralMatrix2(4))
}

func spiralMatrix2(n int) [][]int {
	results := make([][]int, n)
	head := 0
	botom := n - 1
	left := 0
	right := n - 1
	direct := 0
	for i := 0; i < n; i++ {
		results[i] = make([]int, n)
	}
	value := 1
here:
	for value <= n*n {
		switch direct {
		case 0:
			// 右
			i := head
			for j := left; j <= right; j++ {
				results[i][j] = value
				value++
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
				results[i][j] = value
				value++
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
				results[i][j] = value
				value++
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
				results[i][j] = value
				value++
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
