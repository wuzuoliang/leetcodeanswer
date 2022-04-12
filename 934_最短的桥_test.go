package Code

import (
	"math"
	"testing"
)

/**
在给定的二维二进制数组A中，存在两座岛。（岛是由四面相连的 1 形成的一个最大组。）

现在，我们可以将0变为1，以使两座岛连接起来，变成一座岛。

返回必须翻转的0 的最小数目。（可以保证答案至少是 1 。）



示例 1：

输入：A = [[0,1],[1,0]]
输出：1
示例 2：

输入：A = [[0,1,0],[0,0,0],[0,0,1]]
输出：2
示例 3：

输入：A = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
输出：1


提示：

2 <= A.length == A[0].length <= 100
A[i][j] == 0 或 A[i][j] == 1


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shortest-bridge
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// https://www.bilibili.com/video/BV1XU4y1d7E1?p=44
func Test934(t *testing.T) {
	grid := [][]int{{0, 1, 0}, {0, 0, 0}, {1, 0, 0}}
	curs := make([]int, 9)
	record := make([]int, 9)
	t.Log(infect(grid, 0, 1, 3, 3, 0, curs, record))
	t.Log(grid)
	t.Log(curs)
	t.Log(record)
	curs = make([]int, 9)
	record = make([]int, 9)
	t.Log(infect(grid, 2, 0, 3, 3, 0, curs, record))
	t.Log(grid)
	t.Log(curs)
	t.Log(record)

	t.Log(shortestBridge([][]int{{0, 1, 0}, {0, 0, 0}, {1, 0, 0}}))
}

func shortestBridge(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	all := m * n
	curs := make([]int, all)
	next := make([]int, all)
	// 二维变一维
	island := 0
	cover := make([][]int, 2)
	for i := 0; i < 2; i++ {
		cover[i] = make([]int, all)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 { // 当前位置发现了1
				// 把这篇都变成2，抓取这片1组成的初始队列curs,
				// 把这些1到自己的距离都设置成1，存入cover
				queueSize := infect(grid, i, j, m, n, 0, curs, cover[island])
				level := 1
				for queueSize != 0 {
					level++
					// 遍历curs里，cover里距离为0的点
					queueSize = bfs973(n, all, level, queueSize, curs, next, cover[island])
					//fmt.Println("island",island,cover[island])
					curs, next = next, curs
				}
				island++
			}
		}
	}

	min := math.MaxInt
	for i := 0; i < all; i++ {
		min = Min(min, cover[0][i]+cover[1][i])
	}
	return min - 3
}

// index代表队列数量, curs表示当前的下标，record表示下标扩展的层数
func infect(grid [][]int, i, j, m, n, index int, curs, record []int) int {
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != 1 {
		return index
	}
	grid[i][j] = 2

	p := i*m + j
	record[p] = 1
	// 收集到不同的1
	curs[index] = p
	index++
	index = infect(grid, i-1, j, m, n, index, curs, record)
	index = infect(grid, i+1, j, m, n, index, curs, record)
	index = infect(grid, i, j-1, m, n, index, curs, record)
	index = infect(grid, i, j+1, m, n, index, curs, record)
	return index
}

// bfs973 二维原始矩阵
// V 要生成的是第几层 curs V-1层 nexts V 层
// record里获取距离
func bfs973(m, all, level, queueSize int, curs, nexts, record []int) int {
	nextIndex := 0
	for i := 0; i < queueSize; i++ {
		up := -1
		if curs[i] > m {
			up = curs[i] - m
		}
		if up != -1 && record[up] == 0 {
			record[up] = level
			nexts[nextIndex] = up
			nextIndex++
		}

		down := -1
		if curs[i]+m < all {
			down = curs[i] + m
		}
		if down != -1 && record[down] == 0 {
			record[down] = level
			nexts[nextIndex] = down
			nextIndex++
		}
		left := -1
		if curs[i]%m != 0 {
			left = curs[i] - 1
		}
		if left != -1 && record[left] == 0 {
			record[left] = level
			nexts[nextIndex] = left
			nextIndex++
		}

		right := -1
		if curs[i]%m != m-1 {
			right = curs[i] + 1
		}
		if right != -1 && record[right] == 0 {
			record[right] = level
			nexts[nextIndex] = right
			nextIndex++
		}
	}
	//fmt.Println(record)
	return nextIndex
}
