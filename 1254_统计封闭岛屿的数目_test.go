package Code

import "testing"

/**
1、用0表示陆地，用1表示海水。

2、让你计算「封闭岛屿」的数目。所谓「封闭岛屿」就是上下左右全部被1包围的0，也就是说靠边的陆地不算作「封闭岛屿」。
*/
func Test1254(t *testing.T) {
	t.Log(closedIsland([][]int{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}})) //2
}

func closedIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < n; i++ {
		// 把靠上边的岛屿淹掉
		dfs1254(grid, 0, i, m, n)
		// 把靠下边的岛屿淹掉
		dfs1254(grid, m-1, i, m, n)
	}
	for i := 0; i < m; i++ {
		// 把靠左边的岛屿淹掉
		dfs1254(grid, i, 0, m, n)
		// 把靠右边的岛屿淹掉
		dfs1254(grid, i, n-i, m, n)
	}
	// 遍历 grid，剩下的岛屿都是封闭岛屿
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				res++
			}
			dfs1254(grid, i, j, m, n)
		}
	}
	return res
}

func dfs1254(grid [][]int, i, j, m, n int) {
	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}
	if grid[i][j] == 1 {
		return
	}
	grid[i][j] = 1
	dfs1254(grid, i, j-1, m, n)
	dfs1254(grid, i, j+1, m, n)
	dfs1254(grid, i-1, j, m, n)
	dfs1254(grid, i+1, j, m, n)
}
