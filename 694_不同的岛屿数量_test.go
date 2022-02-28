package Code

import (
	"fmt"
	"strings"
	"testing"
)

func Test694(t *testing.T) {
	t.Log(numDistinctIslands([][]int{{1, 1, 0, 1, 1}, {1, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 1}, {1, 1, 0, 1, 1}}))
}

var diffRecords map[string]struct{}

func numDistinctIslands(grid [][]int) int {
	diffRecords = make(map[string]struct{})
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				path := make([]string, 0, m*n)
				dfs694(grid, i, j, m, n, &path, "666")
				pathStr := strings.Join(path, ",")
				fmt.Println(pathStr)
				diffRecords[pathStr] = struct{}{}
			}
		}
	}
	return len(diffRecords)
}

func dfs694(grid [][]int, i, j, m, n int, records *[]string, direct string) {
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	*records = append(*records, direct, "#")
	dfs694(grid, i, j-1, m, n, records, "-1")
	dfs694(grid, i, j+1, m, n, records, "-2")
	dfs694(grid, i-1, j, m, n, records, "-3")
	dfs694(grid, i+1, j, m, n, records, "-4")
	*records = append(*records, "-"+direct, "#")
}
