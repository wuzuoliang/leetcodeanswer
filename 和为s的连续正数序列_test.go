package Code

/**
题目：和为s的连续正数序列
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]
*/

func findContinuousSequence(target int) [][]int {
	result := make([][]int, 0)
	i := 1
	j := 1
	win := 0
	arr := make([]int, target)
	for i := range arr {
		arr[i] = i + 1
	}
	for i <= target/2 {
		if win < target {
			win += j
			j++
		} else if win > target {
			win -= i
			i++
		} else {
			result = append(result, arr[i-1:j-1])
			win -= i
			i++
		}
	}
	return result
}
