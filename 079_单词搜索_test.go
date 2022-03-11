package Code

import (
	"testing"
)

/**
给定一个m x n 二维字符网格board 和一个字符串单词word 。如果word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
输出：true
示例 3：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
输出：false

提示：

m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
board 和 word 仅由大小写英文字母组成


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test79(t *testing.T) {
	//t.Log(exist([][]byte{
	//	{'A','B','C','E'},
	//	{'S','F','C','S'},
	//	{'A','D','E','E'}},"ABCCED"))//true
	//t.Log(exist([][]byte{
	//	{'A','B'},
	//	{'C','D'}},"CDBA"))//true
	t.Log(exist([][]byte{
		{'C', 'A', 'A'},
		{'A', 'A', 'A'},
		{'B', 'C', 'D'}}, "AAB")) //true
}

var existed79 bool

func exist(board [][]byte, word string) bool {
	existed79 = false

	m := len(board)
	n := len(board[0])
	curMatch := 0
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			curMatch = 0
			dfs79(board, word, vis, i, j, &curMatch)
		}
	}
	return existed79
}

func dfs79(board [][]byte, word string, visited [][]bool, i, j int, curMat *int) {
	if *curMat >= len(word) || existed79 == true {
		existed79 = true
		return
	}
	m := len(board)
	n := len(board[0])
	if i >= m || j >= n || i < 0 || j < 0 {
		// 越界了
		return
	}

	if board[i][j] != word[*curMat] {
		return
	}
	if visited[i][j] == true {
		return
	}
	visited[i][j] = true
	*curMat++
	if *curMat == len(word) {
		existed79 = true
		return
	}
	dfs79(board, word, visited, i+1, j, curMat)
	dfs79(board, word, visited, i-1, j, curMat)
	dfs79(board, word, visited, i, j+1, curMat)
	dfs79(board, word, visited, i, j-1, curMat)
	*curMat--
	visited[i][j] = false
}

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func existBest(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] { // 剪枝：当前字符不匹配
			return false
		}
		if k == len(word)-1 { // 单词存在于网格中
			return true
		}
		vis[i][j] = true
		defer func() { vis[i][j] = false }() // 回溯时还原已访问的单元格
		for _, dir := range directions {
			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < h && 0 <= newJ && newJ < w && !vis[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}
