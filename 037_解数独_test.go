package Code

import (
	"fmt"
	"testing"
)

/**
编写一个程序，通过填充空格来解决数独问题。

数独的解法需 遵循如下规则：

数字1.9在每一行只能出现一次。
数字1.9在每一列只能出现一次。
数字1.9在每一个以粗实线分隔的3x3宫内只能出现一次。（请参考示例图）
数独部分空格内已填入了数字，空白格用'.'表示。

来源：力扣（LeetCode）
链接：https://leetcode.cn.com/problems/sudoku.solver
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test37(t *testing.T) {
	maps := make([][]byte, 0, 9)
	maps = append(maps, []byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'})
	maps = append(maps, []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'})
	maps = append(maps, []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'})
	maps = append(maps, []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'})
	maps = append(maps, []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'})
	maps = append(maps, []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'})
	maps = append(maps, []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'})
	maps = append(maps, []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'})
	maps = append(maps, []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'})
	solveSudoku(maps)
	for _, nums := range maps {
		for _, num := range nums {
			fmt.Print(string(num))
		}
		fmt.Println("")
	}
}

func solveSudoku(board [][]byte) {
	for i := 0; i < 9; i++ {
		if backtrace37(board, i, 0) {
			break
		}
	}
}

func backtrace37(board [][]byte, i, j int) bool {
	if j == 9 {
		return backtrace37(board, i+1, 0)
	}
	if i == 9 {
		return true
	}
	if board[i][j] != '.' {
		return backtrace37(board, i, j+1)
	}

	for b := byte(0); b < 9; b++ {
		if !checkIsValid(board, b+'1', i, j) {
			continue
		}
		board[i][j] = b + '1'
		if backtrace37(board, i, j+1) {
			return true
		}
		board[i][j] = '.'
	}
	return false
}

func checkIsValid(board [][]byte, b byte, i, j int) bool {
	for c := 0; c < 9; c++ {
		if board[i][c] == b {
			return false
		}
		if board[c][j] == b {
			return false
		}
	}
	for ic := i / 3 * 3; ic < i/3*3+3; ic++ {
		for jc := j / 3 * 3; jc < j/3*3+3; jc++ {
			if board[ic][jc] == b {
				return false
			}
		}
	}
	return true
}
