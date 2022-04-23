package Code

import "testing"

/**
有一批司机，去A和B的收入表示 [1,2] ,[3,4]
且A和B必须平均分配，求最终整体的最大收益

https://www.bilibili.com/video/BV1334y1v7th?p=6
*/
func TestMaxAllDriversIncome(t *testing.T) {
	t.Log(maxAllDriversIncome([][]int{
		{1, 2}, {3, 4}, {5, 6}, {7, 8},
	}))
}

func maxAllDriversIncome(incomes [][]int) int {
	dirversNum := len(incomes)
	if dirversNum%2 != 0 {
		return 0
	}
	return processDirvers(incomes, 0, dirversNum/2, dirversNum)
}

// index 表示当前的司机
// leftToA 还有多少个司机可以去A
// 总司机数量
// 函数返回从当前index。。最后所有的司机，整体下来平均分配到A和B的最大收入
func processDirvers(incomes [][]int, index int, leftToA int, driversNum int) int {
	if index == driversNum {
		return 0
	}
	if driversNum-index == leftToA {
		return processDirvers(incomes, index+1, leftToA-1, driversNum) + incomes[index][0]
	}
	if leftToA == 0 {
		return processDirvers(incomes, index+1, leftToA, driversNum) + incomes[index][1]
	}
	toA := processDirvers(incomes, index+1, leftToA-1, driversNum) + incomes[index][0]
	toB := processDirvers(incomes, index+1, leftToA, driversNum) + incomes[index][1]
	return Max(toA, toB)
}
