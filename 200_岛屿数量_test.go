package Code

import "testing"

func Test200(t *testing.T) {
	t.Log(numIslands([][]byte{
		{1, 1, 0, 1, 1},
		{1, 0, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 1},
		{1, 1, 0, 1, 1},
	})) // 4
}

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	total := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 发现一个岛屿，数量+1
			if grid[i][j] == 1 {
				total++
			}
			dfs200(grid, i, j, m, n)
		}
	}

	return total
}

func dfs200(grid [][]byte, i, j, m, n int) {
	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}

	if grid[i][j] == 0 {
		return
	}
	// 给岛缩小
	grid[i][j] = 0

	// 淹没上下左右的陆地

	dfs200(grid, i, j-1, m, n)
	dfs200(grid, i, j+1, m, n)
	dfs200(grid, i-1, j, m, n)
	dfs200(grid, i+1, j, m, n)

}
