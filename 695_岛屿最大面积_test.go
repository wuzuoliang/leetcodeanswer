package Code

import "testing"

func Test695(t *testing.T) {}

func maxAresOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	maxArea := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				maxArea = Max(maxArea, dfs695(grid, i, j, m, n))
			}
		}
	}
	return maxArea
}

func dfs695(grid [][]int, i, j, m, n int) int {
	if i < 0 || i >= m || j < 0 || j >= n {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return dfs695(grid, i, j-1, m, n) + dfs695(grid, i, j+1, m, n) + dfs695(grid, i-1, j, m, n) + dfs695(grid, i+1, j, m, n) + 1
}
